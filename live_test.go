// Hand-written: tests against the live API, skipped unless ROICAI_API_KEY is
// set.

package roicai_test

import (
	"context"
	"os"
	"testing"
	"time"

	roicai "github.com/roicai/roicai-go"
	"github.com/roicai/roicai-go/client"
)

func liveClient(t *testing.T) *client.Client {
	t.Helper()
	if os.Getenv("ROICAI_API_KEY") == "" {
		t.Skip("ROICAI_API_KEY is not set")
	}
	return client.NewClient()
}

func TestLiveRetrievesATicker(t *testing.T) {
	c := liveClient(t)

	ticker, err := c.Tickers.Retrieve(context.Background(), &roicai.RetrieveTickerRequest{Identifier: "NASDAQ:AAPL"})
	if err != nil {
		t.Fatal(err)
	}

	if ticker.Symbol != "NASDAQ:AAPL" || ticker.GetName() == nil || *ticker.GetName() != "Apple Inc." {
		t.Errorf("ticker = %+v", ticker)
	}
}

func TestLivePagesThroughAList(t *testing.T) {
	c := liveClient(t)
	ctx := context.Background()

	page, err := c.Tickers.List(ctx, &roicai.ListTickersRequest{Limit: roicai.Int(2)})
	if err != nil {
		t.Fatal(err)
	}
	seen := map[string]bool{}
	for row, err := range page.All(ctx) {
		if err != nil {
			t.Fatal(err)
		}
		seen[row.ID] = true
		if len(seen) == 5 {
			break
		}
	}

	if len(seen) != 5 {
		t.Errorf("got %d distinct tickers, want 5", len(seen))
	}
}

func TestLiveFiltersByDateRange(t *testing.T) {
	c := liveClient(t)
	ctx := context.Background()

	page, err := c.StockPrices.List(ctx, &roicai.ListHistoricalStockPricesRequest{
		Identifier: "NASDAQ:AAPL",
		DateGte:    roicai.String("2025-01-01"),
		DateLte:    roicai.String("2025-01-31"),
		Limit:      roicai.Int(10),
	})
	if err != nil {
		t.Fatal(err)
	}
	from, to := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2025, 1, 31, 0, 0, 0, 0, time.UTC)
	days := 0
	for price, err := range page.All(ctx) {
		if err != nil {
			t.Fatal(err)
		}
		if price.Date.Before(from) || price.Date.After(to) {
			t.Errorf("price on %s is outside January 2025", price.Date.Format(time.DateOnly))
		}
		days++
	}

	if days != 20 {
		t.Errorf("got %d trading days in January 2025, want 20", days)
	}
}
