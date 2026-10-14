# ROIC.ai Go Library

[![Go Reference](https://pkg.go.dev/badge/github.com/roicai/roicai-go.svg)](https://pkg.go.dev/github.com/roicai/roicai-go)
[![Build Status](https://github.com/roicai/roicai-go/actions/workflows/ci.yml/badge.svg)](https://github.com/roicai/roicai-go/actions/workflows/ci.yml)

The ROIC.ai Go library is the official Go client for the [ROIC.ai](https://roic.ai) financial data API and stock API. Use it to get fundamentals for public companies, such as income statements, balance sheets, cash flow statements, financial ratios and valuation multiples, and earnings call transcripts. It also covers historical and latest stock prices, SEC filings and other company documents, the earnings calendar, company profiles and forex rates.

## Documentation

See the [ROIC.ai API reference](https://api.roic.ai/v3.0.0/openapi) for every endpoint, parameter and field, and the [API guides](https://www.roic.ai/api) for authentication, plans and response formats. The [Go package reference](https://pkg.go.dev/github.com/roicai/roicai-go) lists every type and method.

## Requirements

Go 1.23 or later.

## Installation

Add the module to your project with:

```sh
go get github.com/roicai/roicai-go
```

## Usage

The library needs your ROIC.ai API key, which you can find on the [authentication page](https://www.roic.ai/api/docs/authentication) once you sign in. By default the client reads it from the `ROICAI_API_KEY` environment variable:

```go
package main

import (
	"context"
	"fmt"
	"log"

	roicai "github.com/roicai/roicai-go"
	"github.com/roicai/roicai-go/client"
	"github.com/roicai/roicai-go/fundamentals"
	"github.com/roicai/roicai-go/stockprices"
)

func main() {
	c := client.NewClient()
	ctx := context.Background()

	ticker, err := c.Tickers.Retrieve(ctx, &roicai.RetrieveTickerRequest{Identifier: "NASDAQ:AAPL"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*ticker.Name) // Apple Inc.

	statements, err := c.Fundamentals.IncomeStatements.List(ctx, &fundamentals.ListIncomeStatementsRequest{
		Identifier: "NASDAQ:AAPL",
		PeriodType: roicai.PeriodTypeAnnual.Ptr(),
		Limit:      roicai.Int(5),
	})
	if err != nil {
		log.Fatal(err)
	}
	for _, year := range statements.Results {
		fmt.Printf("%d revenue: %.0f\n", *year.FiscalYear, *year.IsSalesRevenueTurnover)
	}

	latest, err := c.StockPrices.Latest.Retrieve(ctx, &stockprices.RetrieveLatestStockPriceRequest{Identifier: "NASDAQ:AAPL"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(latest.Date.Format("2006-01-02"), latest.Close)
}
```

To pass the key yourself, use `option.WithAPIKey`:

```go
c := client.NewClient(option.WithAPIKey(apiKey))
```

Methods are grouped the way the API's URLs are: `/v3.0.0/fundamental/ratios/credit/{identifier}` is `c.Fundamentals.Ratios.Credit.List(...)`. Each group has `List`, `Retrieve` or `Search`. Each method's request type is named after the API operation, such as `ratios.ListCreditRatiosRequest`, and lives in the package of the group above it.

Identify a company by its exchange and symbol separated by a colon (`NASDAQ:AAPL`), its ticker ID (`tkr_...`) or its FIGI.

### Fields and types

Fields are Go names (`PeriodEndDate`) for the API's snake_case JSON (`period_end_date`). Dates and timestamps are `time.Time`.

A field the API can return as `null` is a pointer, which is `nil` when the value is missing. Optional request fields are pointers too. Set them with the helpers `roicai.String`, `roicai.Int`, `roicai.Bool` and `roicai.Float64`, and with `Ptr()` on enums:

```go
page, err := c.Tickers.List(ctx, &roicai.ListTickersRequest{
	Exchange:  roicai.String("NASDAQ"),
	IsPrimary: roicai.Bool(true),
	Status:    roicai.TickerStatusListed.Ptr(),
	Limit:     roicai.Int(100),
})
if err != nil {
	log.Fatal(err)
}
for _, ticker := range page.Results {
	if ticker.Isin != nil {
		fmt.Println(ticker.Symbol, *ticker.Isin)
	}
}
```

## Configuration

### Client options

Pass options to `client.NewClient`:

```go
configured := client.NewClient(
	option.WithAPIKey(apiKey),
	option.WithHTTPClient(&http.Client{Timeout: 30 * time.Second}),
	option.WithHTTPHeader(http.Header{"X-My-App": []string{"reports/1.0"}}),
)
```

| Option | Default | Description |
|---|---|---|
| `option.WithAPIKey` | `ROICAI_API_KEY` | Your API key, sent as `Authorization: Bearer <key>`. |
| `option.WithAPIKeyFunc` | | A function that returns the API key for each request, such as from a secret store. |
| `option.WithBaseURL` | `https://api.roic.ai` | URL that requests are made to. |
| `option.WithHTTPClient` | `http.DefaultClient` | [The `*http.Client` that sends requests.](#using-a-custom-http-client) |
| `option.WithHTTPHeader` | | Headers sent with every request. |
| `option.WithMaxAttempts` | `2` | [How many times to try a request.](#retries) |
| `option.WithoutRetries` | | Turns retries off. |

### Request options

Every method takes the same options as its last arguments, which apply to that request only and take precedence over the client's. A request also takes `option.WithQueryParameters`, for query parameters the library doesn't have a field for:

```go
ticker, err := c.Tickers.Retrieve(
	ctx,
	&roicai.RetrieveTickerRequest{Identifier: "NASDAQ:AAPL"},
	option.WithMaxAttempts(4),
)
```

### Retries

The library retries a request that fails with `408`, `429` or a `5xx` status. By default it makes 2 attempts. It waits for the time in the response's `Retry-After` header, or for a `429` its `X-RateLimit-Reset`, or else backs off exponentially from one second. When the wait would be longer than a minute, such as a `429` once your daily quota runs out, it returns the error at once instead of retrying. Set the number of attempts with `option.WithMaxAttempts`, or turn retries off with `option.WithoutRetries`.

### Configuring timeouts

The library sets no timeout. Use a context to limit one call, retries included:

```go
timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()

ticker, err := c.Tickers.Retrieve(timeoutCtx, &roicai.RetrieveTickerRequest{Identifier: "NASDAQ:AAPL"})
```

A call that times out returns an error that wraps `context.DeadlineExceeded`. To set a timeout for every request, pass an `*http.Client` with a `Timeout`.

### Examining responses

A page from a list method has the status code and headers of its response in `page.StatusCode` and `page.Header`. For other methods, each group has a `WithRawResponse` client whose methods also return them, such as the request ID to quote when you contact support:

```go
raw, err := c.Tickers.WithRawResponse.Retrieve(ctx, &roicai.RetrieveTickerRequest{Identifier: "NASDAQ:AAPL"})
if err != nil {
	log.Fatal(err)
}
fmt.Println(raw.StatusCode, raw.Header.Get("Request-Id"), *raw.Body.Name)
```

### Using a custom HTTP client

Pass your own `*http.Client` to set timeouts, proxies or a transport that logs requests:

```go
type loggingTransport struct{}

func (loggingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	res, err := http.DefaultTransport.RoundTrip(req)
	if err == nil {
		log.Println(req.Method, req.URL, res.StatusCode)
	}
	return res, err
}

logged := client.NewClient(option.WithHTTPClient(&http.Client{Transport: loggingTransport{}}))
```

### Date and fiscal year ranges

Lists of dated rows take `DateGte`, `DateGt`, `DateLte` and `DateLt`, with dates as `YYYY-MM-DD`:

```go
january, err := c.StockPrices.List(ctx, &roicai.ListHistoricalStockPricesRequest{
	Identifier: "NASDAQ:AAPL",
	DateGte:    roicai.String("2025-01-01"),
	DateLte:    roicai.String("2025-01-31"),
	Order:      roicai.OrderAsc.Ptr(),
})
```

They work on stock prices, stock splits, forex prices, earnings calls, the earnings calendar and every fundamentals list. Fundamentals and employee counts also take `FiscalYearGte`, `FiscalYearGt`, `FiscalYearLte` and `FiscalYearLt`:

```go
since2020, err := c.Fundamentals.IncomeStatements.List(ctx, &fundamentals.ListIncomeStatementsRequest{
	Identifier:    "NASDAQ:AAPL",
	PeriodType:    roicai.PeriodTypeAnnual.Ptr(),
	FiscalYearGte: roicai.Int(2020),
})
```

### Auto-pagination

List methods return one page of up to `Limit` rows. You can also have the library fetch the next pages for you as you go. It sends the same request again with the page token from `next_page_url`, so every page uses your filters, and it stops after the last page.

#### Range over `All`

`All` returns an iterator over every item of every page, for use with `range`:

```go
page, err := c.EarningsCalls.List(ctx, &roicai.ListEarningsCallsRequest{
	Identifier: roicai.String("NASDAQ:AAPL"),
	Limit:      roicai.Int(100),
})
if err != nil {
	log.Fatal(err)
}
for call, err := range page.All(ctx) {
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(call.FiscalYear, call.FiscalQuarter)
}
```

Use `break` to stop early. The library doesn't request any page after the one you stopped in. If a request for a page fails, the loop yields its error and ends.

#### `Iterator`

On older code, or to keep the error out of the loop, use `Iterator`:

```go
it := page.Iterator()
for it.Next(ctx) {
	call := it.Current()
	fmt.Println(call.FiscalYear, call.FiscalQuarter)
}
if err := it.Err(); err != nil {
	log.Fatal(err)
}
```

#### Fetching pages yourself

`page.Results` holds the first page. `page.GetNextPage(ctx)` fetches the next one, and returns `core.ErrNoPages` after the last:

```go
for {
	for _, call := range page.Results {
		fmt.Println(call.ID)
	}
	next, err := page.GetNextPage(ctx)
	if errors.Is(err, core.ErrNoPages) {
		break
	}
	if err != nil {
		log.Fatal(err)
	}
	page = next
}
```

The whole response of each page is in `page.Response`, such as `NextPageURL`.

## Handling errors

A request that fails returns an error with the HTTP status and the API's error. Each status the API documents has its own type, such as `*roicai.NotFoundError`, with the error body in `Body`:

```go
_, err := c.Tickers.Retrieve(ctx, &roicai.RetrieveTickerRequest{Identifier: "NASDAQ:NOPE"})

var notFound *roicai.NotFoundError
if errors.As(err, &notFound) {
	// no such ticker
	details := notFound.Body.GetError()
	fmt.Println(details.GetCode(), details.GetMessage(), details.GetRequestID()) // resource_missing ...
}
fmt.Println(err)
// 404: No ticker matches the supplied identifier. (code: resource_missing, request_id: req_...)
```

The API returns `400` (`*roicai.BadRequestError`) for an invalid parameter, `401` (`*roicai.UnauthorizedError`) for a missing or invalid key, `402` (`*roicai.PaymentRequiredError`) when your plan doesn't include the data, `404` (`*roicai.NotFoundError`) for an unknown identifier and `429` (`*roicai.TooManyRequestsError`) when you exceed your rate limit. Every one of them wraps a `*core.APIError`, which has the `StatusCode` and `Header`:

```go
var apiErr *core.APIError
if errors.As(err, &apiErr) {
	fmt.Println(apiErr.StatusCode, apiErr.Header.Get("Request-Id"))
}
```

## Examples

The examples use the client `c` and the context `ctx` from [Usage](#usage).

### Tickers

Retrieve one ticker, search by symbol or company name, or list tickers with filters:

```go
apple, err := c.Tickers.Retrieve(ctx, &roicai.RetrieveTickerRequest{Identifier: "NASDAQ:AAPL"})
if err != nil {
	log.Fatal(err)
}
fmt.Println(*apple.Name, *apple.Isin, *apple.Figi)

bySymbol, err := c.Tickers.Search(ctx, &roicai.SearchTickersRequest{Query: "AAPL", Limit: roicai.Int(5)})
if err != nil {
	log.Fatal(err)
}
byName, err := c.Tickers.Search(ctx, &roicai.SearchTickersRequest{
	Query:    "apple",
	SearchBy: roicai.SearchTickersSearchByName.Ptr(),
	Limit:    roicai.Int(5),
})
if err != nil {
	log.Fatal(err)
}
for _, ticker := range byName.Results {
	fmt.Println(ticker.Symbol)
}

japan, err := c.Tickers.List(ctx, &roicai.ListTickersRequest{
	ListingCountryCode: roicai.String("JP"),
	IsPrimary:          roicai.Bool(true),
})
```

`SearchBy` also matches `cik`, `cusip`, `isin`, `figi` and the other ticker fields exactly.

### Company profile

```go
profile, err := c.Companies.Profiles.Retrieve(ctx, &companies.RetrieveCompanyProfileRequest{Identifier: "NASDAQ:AAPL"})
if err != nil {
	log.Fatal(err)
}
fmt.Println(*profile.Name, *profile.Sector, *profile.Industry, *profile.Ceo)
if next := profile.EarningsSchedule.NextReleaseDate; next != nil {
	fmt.Println("next earnings:", next.Format("2006-01-02"))
}
```

### Stock prices

Historical daily prices come newest first. Prices are unadjusted unless you set `Adjustment` to `splits` or `total_return`:

```go
prices, err := c.StockPrices.List(ctx, &roicai.ListHistoricalStockPricesRequest{
	Identifier: "NASDAQ:AAPL",
	Limit:      roicai.Int(30),
})
if err != nil {
	log.Fatal(err)
}
for _, day := range prices.Results {
	fmt.Println(day.Date.Format("2006-01-02"), *day.Open, *day.High, *day.Low, day.Close, *day.Volume)
}
fmt.Println(*prices.Response.Currency) // USD

splitAdjusted, err := c.StockPrices.List(ctx, &roicai.ListHistoricalStockPricesRequest{
	Identifier: "NASDAQ:AAPL",
	Date:       roicai.String("2020-08-28"),
	Adjustment: roicai.ListHistoricalStockPricesAdjustmentSplits.Ptr(),
})
```

For a date range, see [Date and fiscal year ranges](#date-and-fiscal-year-ranges).

### Financial statements

Income statements, balance sheets and cash flow statements take the same parameters. Set `PeriodType` to `annual`, `quarterly` or `ttm`, and `PeriodLabel` to one fiscal period (`FY`, `Q1` to `Q4`):

```go
annual, err := c.Fundamentals.IncomeStatements.List(ctx, &fundamentals.ListIncomeStatementsRequest{
	Identifier: "NASDAQ:AAPL",
	PeriodType: roicai.PeriodTypeAnnual.Ptr(),
	Limit:      roicai.Int(5),
})
if err != nil {
	log.Fatal(err)
}
for _, year := range annual.Results {
	fmt.Println(*year.FiscalYear, *year.IsSalesRevenueTurnover, *year.IsNetIncome, *year.DilutedEps)
}

quarters, err := c.Fundamentals.BalanceSheets.List(ctx, &fundamentals.ListBalanceSheetsRequest{
	Identifier: "NASDAQ:AAPL",
	PeriodType: roicai.PeriodTypeQuarterly.Ptr(),
	Limit:      roicai.Int(4),
})
if err != nil {
	log.Fatal(err)
}
for _, q := range quarters.Results {
	fmt.Println(*q.FiscalYear, *q.PeriodLabel, *q.BsTotAsset, *q.BsTotLiab)
}

fourthQuarters, err := c.Fundamentals.CashFlowStatements.List(ctx, &fundamentals.ListCashFlowStatementsRequest{
	Identifier:  "NASDAQ:AAPL",
	PeriodType:  roicai.PeriodTypeQuarterly.Ptr(),
	PeriodLabel: roicai.PeriodLabelQ4.Ptr(),
	Limit:       roicai.Int(4),
})
if err != nil {
	log.Fatal(err)
}
for _, q := range fourthQuarters.Results {
	fmt.Println(*q.FiscalYear, *q.CfCashFromOper, *q.CfFreeCashFlow)
}
```

Rows come newest first. Amounts are in `Currency`, and line items use short names: `IsSalesRevenueTurnover` is revenue, `BsTotAsset` is total assets and `CfCapExpenditures` is capital expenditures. `c.Fundamentals.PerShareMetrics` and `c.Fundamentals.EnterpriseValues` work the same way.

### Ratios and multiples

Ratios are grouped into `Profitability`, `Credit`, `Liquidity`, `WorkingCapital` and `YieldAnalysis`. Margins and returns are percentages (`46.9` is 46.9%):

```go
profitability, err := c.Fundamentals.Ratios.Profitability.List(ctx, &ratios.ListProfitabilityRatiosRequest{
	Identifier: "NASDAQ:AAPL",
	PeriodType: roicai.PeriodTypeAnnual.Ptr(),
	Limit:      roicai.Int(5),
})
if err != nil {
	log.Fatal(err)
}
for _, year := range profitability.Results {
	fmt.Println(*year.FiscalYear, *year.GrossMargin, *year.OperMargin, *year.ReturnOnInvCapital)
}

credit, err := c.Fundamentals.Ratios.Credit.List(ctx, &ratios.ListCreditRatiosRequest{
	Identifier: "NASDAQ:AAPL",
	Limit:      roicai.Int(1),
})
if err != nil {
	log.Fatal(err)
}
fmt.Println(*credit.Results[0].NetDebtToEbitda)

multiples, err := c.Fundamentals.Multiples.List(ctx, &fundamentals.ListValuationMultiplesRequest{
	Identifier: "NASDAQ:AAPL",
	PeriodType: roicai.PeriodTypeQuarterly.Ptr(),
	Limit:      roicai.Int(1),
})
if err != nil {
	log.Fatal(err)
}
last := multiples.Results[0]
fmt.Println(*last.PeRatio, *last.EvToTtmEbitda, *last.PrToBookRatio, *last.PrToFreeCashFlow)
```

### Earnings call transcripts

List a company's earnings calls, then retrieve a transcript by fiscal year and quarter. Free plans can access each company's latest two calls:

```go
calls, err := c.EarningsCalls.List(ctx, &roicai.ListEarningsCallsRequest{
	Identifier: roicai.String("NASDAQ:AAPL"),
	Limit:      roicai.Int(8),
})
if err != nil {
	log.Fatal(err)
}
for _, call := range calls.Results {
	fmt.Println(call.FiscalYear, call.FiscalQuarter, call.Date)
}

call, err := c.EarningsCalls.Retrieve(ctx, &roicai.RetrieveEarningsCallRequest{
	Identifier:    "NASDAQ:AAPL",
	FiscalYear:    2026,
	FiscalQuarter: 3,
})
if err != nil {
	log.Fatal(err)
}
for _, turn := range call.EarningsCallJSON.Transcript {
	if turn.Speaker != nil {
		fmt.Printf("%s: %s\n", *turn.Speaker, turn.Text)
	}
}

text, err := c.EarningsCalls.Retrieve(ctx, &roicai.RetrieveEarningsCallRequest{
	Identifier:    "NASDAQ:AAPL",
	FiscalYear:    2026,
	FiscalQuarter: 3,
	Format:        roicai.RetrieveEarningsCallFormatText.Ptr(),
})
if err != nil {
	log.Fatal(err)
}
fmt.Println(len(text.EarningsCallText.Transcript))
```

By default the transcript is a list of turns, in `EarningsCallJSON`, each with a `Speaker` and `Text`. With the `text` format it is one string, in `EarningsCallText`.

Leave out `Identifier` to list calls from every company, and filter by `Date`, `FiscalYear` or `FiscalQuarter`:

```go
sameDay, err := c.EarningsCalls.List(ctx, &roicai.ListEarningsCallsRequest{Date: roicai.String("2026-07-30")})
```

### Earnings calendar

Each event has the release date and session, and the EPS and revenue estimates and actuals:

```go
schedule, err := c.Calendar.Earnings.List(ctx, &calendar.ListEarningsCalendarRequest{
	Identifier: roicai.String("NASDAQ:AAPL"),
	Limit:      roicai.Int(4),
})
if err != nil {
	log.Fatal(err)
}
for _, event := range schedule.Results {
	fmt.Print(*event.FiscalPeriod, " ", event.ReleaseDate.Format("2006-01-02"))
	if event.EpsEstimate != nil && event.EpsActual != nil {
		fmt.Print(" EPS ", *event.EpsActual, " vs ", *event.EpsEstimate, " expected")
	}
	fmt.Println()
}

oneDay, err := c.Calendar.Earnings.List(ctx, &calendar.ListEarningsCalendarRequest{Date: roicai.String("2026-10-28")})
```

`EpsEstimate` and `EpsActual` are `nil` until there's an estimate or the company has reported.

### SEC filings and company documents

A company event is one reporting period, with the IDs of its transcripts, call audio, slides and filings, such as 10-Q and 8-K reports. Retrieve a document by ID to get a signed download URL, which needs no API key:

```go
events, err := c.CompanyEvents.List(ctx, &roicai.ListCompanyEventsRequest{
	Identifier: roicai.String("NASDAQ:AAPL"),
	Limit:      roicai.Int(4),
})
if err != nil {
	log.Fatal(err)
}
for _, event := range events.Results {
	for _, filing := range event.Filings {
		if filing.Form == nil || *filing.Form != "10-Q" {
			continue
		}
		download, err := c.Documents.Filings.Retrieve(ctx, &documents.RetrieveFilingRequest{ID: filing.ID})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(download.Filename, download.URL, download.ExpiresAt)
	}
}
```

`c.Documents.Slides`, `Audio`, `Transcripts.Edited` and `Transcripts.Raw` retrieve the other kinds by ID.

### Forex

Identify a currency pair as `FX:` followed by the base and quote currencies:

```go
pairs, err := c.ForexPairs.List(ctx, &roicai.ListForexPairsRequest{
	BaseCurrency:  roicai.String("EUR"),
	QuoteCurrency: roicai.String("USD"),
})
if err != nil {
	log.Fatal(err)
}
fmt.Println(len(pairs.Results))

eurusd, err := c.ForexPrices.List(ctx, &roicai.ListHistoricalForexPricesRequest{
	Identifier: "FX:EURUSD",
	Limit:      roicai.Int(30),
})
if err != nil {
	log.Fatal(err)
}
for _, day := range eurusd.Results {
	fmt.Println(day.Date.Format("2006-01-02"), day.Close)
}

rate, err := c.ForexPrices.Latest.Retrieve(ctx, &forexprices.RetrieveLatestForexPriceRequest{Identifier: "FX:EURUSD"})
if err != nil {
	log.Fatal(err)
}
fmt.Println(rate.Date.Format("2006-01-02"), rate.Close)
```

## Support

Report bugs and request features in [GitHub issues](https://github.com/roicai/roicai-go/issues). For questions about your account, plan or the data, contact ROIC.ai through [roic.ai](https://roic.ai).

The library follows semver. While it is `v0`, a minor version can include breaking changes, so pin the minor version you use.

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) to build and test the library, regenerate it from the API, and release it.

## License

[MIT](LICENSE)
