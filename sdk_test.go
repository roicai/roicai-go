// Hand-written: offline tests of the SDK against a local server.

package roicai_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	roicai "github.com/roicai/roicai-go"
	"github.com/roicai/roicai-go/client"
	"github.com/roicai/roicai-go/core"
	"github.com/roicai/roicai-go/option"
)

// server answers each request with the next of responses, as JSON with its
// status, and records the requests.
type server struct {
	*httptest.Server
	mu       sync.Mutex
	requests []*http.Request
}

type response struct {
	status int
	body   any
}

func newServer(t *testing.T, responses ...response) *server {
	t.Helper()
	s := &server{}
	s.Server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.mu.Lock()
		i := len(s.requests)
		s.requests = append(s.requests, r)
		s.mu.Unlock()
		if i >= len(responses) {
			t.Errorf("unexpected request %d: %s", i+1, r.URL)
			w.WriteHeader(http.StatusTeapot)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(responses[i].status)
		json.NewEncoder(w).Encode(responses[i].body)
	}))
	t.Cleanup(s.Close)
	return s
}

func (s *server) client(opts ...option.RequestOption) *client.Client {
	return client.NewClient(append([]option.RequestOption{
		option.WithBaseURL(s.URL),
		option.WithAPIKey("test-key"),
		option.WithoutRetries(),
	}, opts...)...)
}

func ticker(symbol string) map[string]any {
	return map[string]any{"id": "tkr_" + symbol, "object": "ticker", "symbol": symbol, "exchange": "NASDAQ"}
}

func tickerPage(next string, symbols ...string) map[string]any {
	data := []any{}
	for _, symbol := range symbols {
		data = append(data, ticker(symbol))
	}
	page := map[string]any{"object": "list", "data": data, "next_page_url": nil}
	if next != "" {
		page["next_page_url"] = next
	}
	return page
}

func TestDefaultBaseURL(t *testing.T) {
	if roicai.Environments.Default != "https://api.roic.ai" {
		t.Fatalf("Environments.Default = %q", roicai.Environments.Default)
	}
}

func TestSendsAPIKeyAndUserAgent(t *testing.T) {
	s := newServer(t, response{200, ticker("NASDAQ:AAPL")}, response{200, ticker("NASDAQ:AAPL")})
	ctx := context.Background()

	if _, err := s.client().Tickers.Retrieve(ctx, &roicai.RetrieveTickerRequest{Identifier: "NASDAQ:AAPL"}); err != nil {
		t.Fatal(err)
	}
	t.Setenv("ROICAI_API_KEY", "env-key")
	fromEnv := client.NewClient(option.WithBaseURL(s.URL), option.WithoutRetries())
	if _, err := fromEnv.Tickers.Retrieve(ctx, &roicai.RetrieveTickerRequest{Identifier: "NASDAQ:AAPL"}); err != nil {
		t.Fatal(err)
	}

	if got := s.requests[0].Header.Get("Authorization"); got != "Bearer test-key" {
		t.Errorf("Authorization = %q, want Bearer test-key", got)
	}
	if got := s.requests[1].Header.Get("Authorization"); got != "Bearer env-key" {
		t.Errorf("Authorization from ROICAI_API_KEY = %q, want Bearer env-key", got)
	}
	if got, want := s.requests[0].Header.Get("User-Agent"), "roicai-go/"+roicai.Version; got != want {
		t.Errorf("User-Agent = %q, want %q", got, want)
	}
	if got := s.requests[0].URL.Path; got != "/v3.0.0/tickers/NASDAQ:AAPL" {
		t.Errorf("path = %q", got)
	}
}

func TestDecodesFields(t *testing.T) {
	s := newServer(t, response{200, map[string]any{
		"object": "list",
		"data": []any{map[string]any{
			"symbol": "NASDAQ:AAPL", "date": "2025-01-31", "close": 236, "open": nil, "volume": 1000,
		}},
		"next_page_url": nil,
	}})

	page, err := s.client().StockPrices.List(context.Background(), &roicai.ListHistoricalStockPricesRequest{Identifier: "NASDAQ:AAPL"})
	if err != nil {
		t.Fatal(err)
	}

	price := page.Results[0]
	if want := time.Date(2025, 1, 31, 0, 0, 0, 0, time.UTC); !price.Date.Equal(want) {
		t.Errorf("Date = %v, want %v", price.Date, want)
	}
	if price.Close != 236 || price.Open != nil || *price.Volume != 1000 {
		t.Errorf("price = %+v", price)
	}
}

func TestSendsRangeFiltersUnderTheirAPINames(t *testing.T) {
	s := newServer(t, response{200, map[string]any{"object": "list", "data": []any{}, "next_page_url": nil}})

	_, err := s.client().StockPrices.List(context.Background(), &roicai.ListHistoricalStockPricesRequest{
		Identifier: "NASDAQ:AAPL",
		DateGte:    roicai.String("2025-01-01"),
		DateLte:    roicai.String("2025-01-31"),
		Limit:      roicai.Int(5),
	})
	if err != nil {
		t.Fatal(err)
	}

	query := s.requests[0].URL.Query()
	if query.Get("date.gte") != "2025-01-01" || query.Get("date.lte") != "2025-01-31" || query.Get("limit") != "5" {
		t.Errorf("query = %s", s.requests[0].URL.RawQuery)
	}
}

func TestAllFetchesEveryPageWithTheSameFilters(t *testing.T) {
	s := newServer(t,
		response{200, tickerPage("/v3.0.0/tickers?limit=2&page=tok2", "A", "B")},
		response{200, tickerPage("/v3.0.0/tickers?page=tok3", "C", "D")},
		response{200, tickerPage("", "E")},
	)
	ctx := context.Background()

	page, err := s.client().Tickers.List(ctx, &roicai.ListTickersRequest{Exchange: roicai.String("NASDAQ"), Limit: roicai.Int(2)})
	if err != nil {
		t.Fatal(err)
	}
	var symbols []string
	for row, err := range page.All(ctx) {
		if err != nil {
			t.Fatal(err)
		}
		symbols = append(symbols, row.Symbol)
	}

	if got := strings.Join(symbols, ","); got != "A,B,C,D,E" {
		t.Errorf("symbols = %s, want A,B,C,D,E", got)
	}
	if len(s.requests) != 3 {
		t.Fatalf("made %d requests, want 3", len(s.requests))
	}
	for i, wantPage := range []string{"", "tok2", "tok3"} {
		query := s.requests[i].URL.Query()
		if query.Get("page") != wantPage || query.Get("exchange") != "NASDAQ" || query.Get("limit") != "2" {
			t.Errorf("request %d query = %s", i+1, s.requests[i].URL.RawQuery)
		}
	}
}

func TestBreakingOutOfAllStopsFetching(t *testing.T) {
	s := newServer(t,
		response{200, tickerPage("/v3.0.0/tickers?page=tok2", "A", "B")},
		response{200, tickerPage("/v3.0.0/tickers?page=tok3", "C", "D")},
	)
	ctx := context.Background()

	page, err := s.client().Tickers.List(ctx, &roicai.ListTickersRequest{})
	if err != nil {
		t.Fatal(err)
	}
	n := 0
	for _, err := range page.All(ctx) {
		if err != nil {
			t.Fatal(err)
		}
		if n++; n == 3 {
			break
		}
	}

	if len(s.requests) != 2 {
		t.Errorf("made %d requests, want 2", len(s.requests))
	}
}

func TestAllEndsWithTheErrorOfAFailedPage(t *testing.T) {
	s := newServer(t,
		response{200, tickerPage("/v3.0.0/tickers?page=tok2", "A")},
		response{500, map[string]any{"error": map[string]any{"code": "internal_error", "message": "Something went wrong."}}},
	)
	ctx := context.Background()

	page, err := s.client().Tickers.List(ctx, &roicai.ListTickersRequest{})
	if err != nil {
		t.Fatal(err)
	}
	var rows int
	var last error
	for _, err := range page.All(ctx) {
		if err != nil {
			last = err
			continue
		}
		rows++
	}

	var serverError *roicai.InternalServerError
	if rows != 1 || !errors.As(last, &serverError) {
		t.Errorf("rows = %d, err = %v; want 1 row and an InternalServerError", rows, last)
	}
}

func TestErrors(t *testing.T) {
	s := newServer(t, response{404, map[string]any{"error": map[string]any{
		"type": "invalid_request_error", "code": "resource_missing", "param": "identifier",
		"message": "No ticker matches the supplied identifier.", "request_id": "req_123",
	}}})

	_, err := s.client().Tickers.Retrieve(context.Background(), &roicai.RetrieveTickerRequest{Identifier: "NASDAQ:NOPE"})

	var notFound *roicai.NotFoundError
	if !errors.As(err, &notFound) {
		t.Fatalf("err = %#v, want a *roicai.NotFoundError", err)
	}
	if notFound.StatusCode != 404 || notFound.Body.GetError().GetCode() != "resource_missing" {
		t.Errorf("StatusCode = %d, code = %q", notFound.StatusCode, notFound.Body.GetError().GetCode())
	}
	want := "404: No ticker matches the supplied identifier. (code: resource_missing, request_id: req_123)"
	if err.Error() != want {
		t.Errorf("Error() = %q, want %q", err.Error(), want)
	}
}

func TestFetchesPagesOfOneListConcurrently(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tickerPage("/v3.0.0/tickers?page=next", "A"))
	}))
	t.Cleanup(s.Close)
	c := client.NewClient(option.WithBaseURL(s.URL), option.WithAPIKey("k"), option.WithoutRetries())
	ctx := context.Background()

	first, err := c.Tickers.List(ctx, &roicai.ListTickersRequest{Exchange: roicai.String("NASDAQ")})
	if err != nil {
		t.Fatal(err)
	}
	var wg sync.WaitGroup
	for range 20 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if _, err := first.GetNextPage(ctx); err != nil {
				t.Error(err)
			}
		}()
	}
	wg.Wait()
}

func TestRejectsAnEmptyPathParameter(t *testing.T) {
	s := newServer(t)

	_, err := s.client().Tickers.Retrieve(context.Background(), &roicai.RetrieveTickerRequest{Identifier: ""})

	if err == nil || !strings.Contains(err.Error(), "path parameter") {
		t.Errorf("err = %v, want an empty path parameter error", err)
	}
	if len(s.requests) != 0 {
		t.Errorf("made %d requests, want none", len(s.requests))
	}
}

func TestTrimsTrailingSlashesFromTheBaseURL(t *testing.T) {
	s := newServer(t, response{200, ticker("NASDAQ:AAPL")})

	_, err := s.client(option.WithBaseURL(s.URL+"/")).Tickers.Retrieve(context.Background(), &roicai.RetrieveTickerRequest{Identifier: "NASDAQ:AAPL"})
	if err != nil {
		t.Fatal(err)
	}

	if got := s.requests[0].URL.Path; got != "/v3.0.0/tickers/NASDAQ:AAPL" {
		t.Errorf("path = %q", got)
	}
}

func TestDoesNotRetryWhenTheServerAsksForALongWait(t *testing.T) {
	var requests int
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests++
		w.Header().Set("Retry-After", "36000")
		w.WriteHeader(http.StatusTooManyRequests)
	}))
	t.Cleanup(s.Close)
	c := client.NewClient(option.WithBaseURL(s.URL), option.WithAPIKey("k"), option.WithMaxAttempts(3))
	start := time.Now()

	_, err := c.Tickers.Retrieve(context.Background(), &roicai.RetrieveTickerRequest{Identifier: "NASDAQ:AAPL"})

	var tooMany *roicai.TooManyRequestsError
	if !errors.As(err, &tooMany) {
		t.Errorf("err = %v, want a *roicai.TooManyRequestsError", err)
	}
	if requests != 1 || time.Since(start) > 5*time.Second {
		t.Errorf("made %d requests in %s, want 1 at once", requests, time.Since(start))
	}
}

func TestRetriesAServerErrorDespiteAFarRateLimitReset(t *testing.T) {
	var requests int
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests++
		w.Header().Set("X-RateLimit-Reset", strconv.FormatInt(time.Now().Add(10*time.Hour).Unix(), 10))
		if requests == 1 {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ticker("NASDAQ:AAPL"))
	}))
	t.Cleanup(s.Close)
	c := client.NewClient(option.WithBaseURL(s.URL), option.WithAPIKey("k"))
	start := time.Now()

	_, err := c.Tickers.Retrieve(context.Background(), &roicai.RetrieveTickerRequest{Identifier: "NASDAQ:AAPL"})

	if err != nil || requests != 2 {
		t.Errorf("err = %v after %d requests, want success on the second", err, requests)
	}
	if elapsed := time.Since(start); elapsed > 5*time.Second {
		t.Errorf("took %s, want the one-second backoff rather than the rate-limit reset", elapsed)
	}
}

func TestAPIKeyFuncTakesPrecedenceOverTheEnvironment(t *testing.T) {
	t.Setenv("ROICAI_API_KEY", "env-key")
	s := newServer(t, response{200, ticker("NASDAQ:AAPL")})
	c := client.NewClient(option.WithBaseURL(s.URL), option.WithAPIKeyFunc(func() (string, error) { return "func-key", nil }))

	if _, err := c.Tickers.Retrieve(context.Background(), &roicai.RetrieveTickerRequest{Identifier: "NASDAQ:AAPL"}); err != nil {
		t.Fatal(err)
	}

	if got := s.requests[0].Header.Get("Authorization"); got != "Bearer func-key" {
		t.Errorf("Authorization = %q, want Bearer func-key", got)
	}
}

func TestErrorsWithoutTheAPIsErrorBody(t *testing.T) {
	for _, tc := range []struct {
		name, contentType, body string
		status                  int
		want                    string
		as                      func(error) bool
	}{
		{"empty 404", "", "", 404, "404: Not Found", func(err error) bool { var e *roicai.NotFoundError; return errors.As(err, &e) && e.Body == nil }},
		{"HTML 503", "text/html", "<html><body><h1>503 Service Temporarily Unavailable</h1></body></html>", 503, "503: Service Unavailable", func(err error) bool { var e *roicai.ServiceUnavailableError; return errors.As(err, &e) }},
		{"plain 502", "text/plain", "upstream connect error", 502, "502: Bad Gateway: upstream connect error", func(err error) bool { var e *core.APIError; return errors.As(err, &e) && e.StatusCode == 502 }},
	} {
		t.Run(tc.name, func(t *testing.T) {
			s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if tc.contentType != "" {
					w.Header().Set("Content-Type", tc.contentType)
				}
				w.WriteHeader(tc.status)
				w.Write([]byte(tc.body))
			}))
			t.Cleanup(s.Close)
			c := client.NewClient(option.WithBaseURL(s.URL), option.WithAPIKey("k"), option.WithoutRetries())

			_, err := c.Tickers.Retrieve(context.Background(), &roicai.RetrieveTickerRequest{Identifier: "NASDAQ:AAPL"})

			if err == nil || err.Error() != tc.want {
				t.Errorf("Error() = %v, want %q", err, tc.want)
			}
			if !tc.as(err) {
				t.Errorf("err = %#v doesn't have the expected type", err)
			}
		})
	}
}
