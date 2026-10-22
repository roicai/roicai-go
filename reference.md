# Reference
## Exchanges
<details><summary><code>client.Exchanges.List() -> *roicai.ExchangeList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns every exchange in one response, with segment MICs, names, markets, and symbol counts. Each code can be passed to the ticker-list exchange filter.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Exchanges.List(
    context.TODO(),
)
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Markets
<details><summary><code>client.Markets.List() -> *roicai.MarketList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns every market in one response, with its exchanges, listing country, and symbol count. Each code can be passed to the ticker-search market filter.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Markets.List(
    context.TODO(),
)
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Sectors
<details><summary><code>client.Sectors.List() -> *roicai.SectorList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns every company sector in one response, with symbol counts.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Sectors.List(
    context.TODO(),
)
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Industries
<details><summary><code>client.Industries.List() -> *roicai.IndustryList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns every company industry in one response, with symbol counts.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Industries.List(
    context.TODO(),
)
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Countries
<details><summary><code>client.Countries.List() -> *roicai.CountryList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns every company country in one response, with names, ISO codes, and symbol counts.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Countries.List(
    context.TODO(),
)
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Tickers
<details><summary><code>client.Tickers.List() -> *roicai.TickerList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a filtered page of ticker resources.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &roicai.ListTickersRequest{}
client.Tickers.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**exchange:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*string` — Filter by exact security type: 'dr', 'fund', or 'stock'.
    
</dd>
</dl>

<dl>
<dd>

**isPrimary:** `*bool` 
    
</dd>
</dl>

<dl>
<dd>

**status:** `*roicai.TickerStatus` 
    
</dd>
</dl>

<dl>
<dd>

**listingCountryCode:** `*string` — Filter by two-letter listing country code: the country of the exchange where the ticker trades, not the company's domicile, e.g. 'US' or 'JP'.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*roicai.Order` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Maximum number of resources to return. Defaults to 500.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*string` — Opaque page token from a previous response. Prefer following the returned page URLs directly.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tickers.Search() -> *roicai.TickerList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Ranks ticker codes and qualified symbols with typo-tolerant ranking by default. Set 'search_by' to 'name' to rank company names, or to another ticker field (such as 'isin' or 'figi') to match that field exactly.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &roicai.SearchTickersRequest{
    Query: "query",
}
client.Tickers.Search(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**query:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**searchBy:** `*roicai.SearchTickersSearchBy` — Fields used for matching. 'symbol' (the default) ranks ticker codes and qualified symbols; 'name' ranks company names. Any other value matches that exact ticker field.
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Maximum number of resources to return. Defaults to 500.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*string` — Opaque page token from a previous response. Prefer following the returned page URLs directly.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Tickers.Retrieve(Identifier) -> *roicai.Ticker</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves one ticker by public ID or supported market identifier.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &roicai.RetrieveTickerRequest{
    Identifier: "NASDAQ:AAPL",
}
client.Tickers.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**identifier:** `string` — A ticker ID, an exchange and symbol separated by a colon (e.g. `NASDAQ:AAPL`), or a FIGI.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## StockSplits
<details><summary><code>client.StockSplits.List() -> *roicai.StockSplitList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a page of per-ticker stock-split events across all tickers, or one ticker via the identifier filter.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &roicai.ListStockSplitsRequest{
    Identifier: roicai.String(
        "NASDAQ:AAPL",
    ),
}
client.StockSplits.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**identifier:** `*string` — Restrict results to a single ticker. Accepts an exchange and symbol separated by a colon (e.g. `NASDAQ:AAPL`), a ticker ID, or a FIGI. Stock splits are per ticker, so results are the split events for that exact listing. Omit to list splits across all tickers.
    
</dd>
</dl>

<dl>
<dd>

**date:** `*string` — Return splits executed on exactly this date.
    
</dd>
</dl>

<dl>
<dd>

**dateGte:** `*string` — Return splits whose execution date is on or after this date.
    
</dd>
</dl>

<dl>
<dd>

**dateGt:** `*string` — Return splits whose execution date is strictly after this date.
    
</dd>
</dl>

<dl>
<dd>

**dateLte:** `*string` — Return splits whose execution date is on or before this date.
    
</dd>
</dl>

<dl>
<dd>

**dateLt:** `*string` — Return splits whose execution date is strictly before this date.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*roicai.Order` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Maximum number of resources to return. Defaults to 100.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*string` — Opaque page token from a previous response. Prefer following the returned page URLs directly.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## StockPrices
<details><summary><code>client.StockPrices.List(Identifier) -> *roicai.HistoricalPriceList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a page of historical daily prices for a ticker, symbol, or other supported security identifier.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &roicai.ListHistoricalStockPricesRequest{
    Identifier: "NASDAQ:AAPL",
}
client.StockPrices.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**identifier:** `string` — A ticker ID, an exchange and symbol separated by a colon (e.g. `NASDAQ:AAPL`), or a FIGI.
    
</dd>
</dl>

<dl>
<dd>

**date:** `*string` — Return prices from exactly this date.
    
</dd>
</dl>

<dl>
<dd>

**dateGte:** `*string` — Return prices whose date is on or after this date.
    
</dd>
</dl>

<dl>
<dd>

**dateGt:** `*string` — Return prices whose date is strictly after this date.
    
</dd>
</dl>

<dl>
<dd>

**dateLte:** `*string` — Return prices whose date is on or before this date.
    
</dd>
</dl>

<dl>
<dd>

**dateLt:** `*string` — Return prices whose date is strictly before this date.
    
</dd>
</dl>

<dl>
<dd>

**adjustment:** `*roicai.ListHistoricalStockPricesAdjustment` 
    
</dd>
</dl>

<dl>
<dd>

**adjusted:** `*roicai.ListHistoricalStockPricesAdjusted` 
    
</dd>
</dl>

<dl>
<dd>

**order:** `*roicai.Order` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Maximum number of resources to return. Defaults to 250.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*string` — Opaque page token from a previous response. Prefer following the returned page URLs directly.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## ForexPairs
<details><summary><code>client.ForexPairs.List() -> *roicai.ForexPairList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a page of tradable currency pairs with their currency codes and human-readable names. Use a pair `symbol` from this list with the forex price endpoints.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &roicai.ListForexPairsRequest{
    BaseCurrency: roicai.String(
        "EUR",
    ),
    QuoteCurrency: roicai.String(
        "EUR",
    ),
    Currency: roicai.String(
        "EUR",
    ),
    Search: roicai.String(
        "yen",
    ),
}
client.ForexPairs.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**baseCurrency:** `*string` — Return only pairs quoted from this base currency.
    
</dd>
</dl>

<dl>
<dd>

**quoteCurrency:** `*string` — Return only pairs quoted into this quote currency.
    
</dd>
</dl>

<dl>
<dd>

**currency:** `*string` — Return every pair in which this currency appears, on either side.
    
</dd>
</dl>

<dl>
<dd>

**search:** `*string` — Case-insensitive substring match against the pair symbol, both currency codes, and the human-readable name.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*roicai.Order` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Maximum number of resources to return. Defaults to 500.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*string` — Opaque page token from a previous response. Prefer following the returned page URLs directly.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## ForexPrices
<details><summary><code>client.ForexPrices.List(Identifier) -> *roicai.ForexPriceList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a page of historical daily prices for a currency pair, addressed by bare currency-pair code or qualified pair symbol.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &roicai.ListHistoricalForexPricesRequest{
    Identifier: "EURUSD",
}
client.ForexPrices.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**identifier:** `string` — A currency pair as `BASEQUOTE` or `BASE/QUOTE`, or a qualified pair symbol (e.g. `FX:EURUSD`). The `symbol` of every `/forex-pairs` result is accepted here.
    
</dd>
</dl>

<dl>
<dd>

**date:** `*string` — Return prices from exactly this date.
    
</dd>
</dl>

<dl>
<dd>

**dateGte:** `*string` — Return prices whose date is on or after this date.
    
</dd>
</dl>

<dl>
<dd>

**dateGt:** `*string` — Return prices whose date is strictly after this date.
    
</dd>
</dl>

<dl>
<dd>

**dateLte:** `*string` — Return prices whose date is on or before this date.
    
</dd>
</dl>

<dl>
<dd>

**dateLt:** `*string` — Return prices whose date is strictly before this date.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*roicai.Order` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Maximum number of resources to return. Defaults to 250.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*string` — Opaque page token from a previous response. Prefer following the returned page URLs directly.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## EarningsCalls
<details><summary><code>client.EarningsCalls.List() -> *roicai.EarningsCallList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns earnings-call summaries across all companies. Free plans can access each company’s latest two calls; paid-plan history follows the plan’s lookback limit. Filter by company identifier, fiscal period, or a call-date range, then retrieve a full transcript by company identifier and fiscal period.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &roicai.ListEarningsCallsRequest{
    Identifier: roicai.String(
        "NASDAQ:AAPL",
    ),
}
client.EarningsCalls.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**identifier:** `*string` — Restrict results to a single company. Accepts an exchange and symbol separated by a colon (e.g. `NASDAQ:AAPL`), a ticker ID, or a FIGI; the response `symbol` echoes the listing you requested. Omit to list earnings calls across all companies.
    
</dd>
</dl>

<dl>
<dd>

**fiscalYear:** `*int` — Return only calls that report on this fiscal year, within the caller’s plan entitlement.
    
</dd>
</dl>

<dl>
<dd>

**fiscalQuarter:** `*int` — Return only calls that report on this fiscal quarter, from 1 to 4, within the caller’s plan entitlement.
    
</dd>
</dl>

<dl>
<dd>

**date:** `*string` — Return calls that occurred on this UTC calendar date.
    
</dd>
</dl>

<dl>
<dd>

**dateGte:** `*string` — Return calls on or after this UTC calendar date.
    
</dd>
</dl>

<dl>
<dd>

**dateGt:** `*string` — Return calls after this UTC calendar date.
    
</dd>
</dl>

<dl>
<dd>

**dateLte:** `*string` — Return calls on or before this UTC calendar date.
    
</dd>
</dl>

<dl>
<dd>

**dateLt:** `*string` — Return calls before this UTC calendar date.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*roicai.Order` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Maximum number of resources to return. Defaults to 50.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*string` — Opaque page token from a previous response. Prefer following the returned page URLs directly.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.EarningsCalls.Retrieve(Identifier) -> *roicai.EarningsCall</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns one complete earnings-call transcript for a company and fiscal period. Free plans can retrieve the company’s latest two calls; paid-plan history follows the plan’s lookback limit. Identify the company by an exchange and symbol separated by a colon (e.g. `NASDAQ:AAPL`), a ticker ID, or a FIGI, and select the call with the `fiscal_year` and `fiscal_quarter` query parameters.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &roicai.RetrieveEarningsCallRequest{
    Identifier: "NASDAQ:AAPL",
    FiscalYear: 1,
    FiscalQuarter: 1,
}
client.EarningsCalls.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**identifier:** `string` — The company whose transcript to retrieve. Accepts an exchange and symbol separated by a colon (e.g. `NASDAQ:AAPL`), a ticker ID, or a FIGI.
    
</dd>
</dl>

<dl>
<dd>

**fiscalYear:** `int` — The fiscal year the call reports on, within the caller’s plan entitlement.
    
</dd>
</dl>

<dl>
<dd>

**fiscalQuarter:** `int` — The fiscal quarter the call reports on, from 1 to 4, within the caller’s plan entitlement.
    
</dd>
</dl>

<dl>
<dd>

**format:** `*roicai.RetrieveEarningsCallFormat` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Company Events
<details><summary><code>client.CompanyEvents.List() -> *roicai.CompanyEventList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns earnings events with everything attached to them: the transcript in `transcripts`, one entry per edition (`raw` and `edited`) with a stable id, audio recordings, slide decks, and filings. Only events with at least one of those are listed, newest fiscal period first. Free plans see each company’s latest two events; Individual plans follow the plan’s lookback; Professional and Enterprise plans see everything. Filter by company identifier or by content kind.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &roicai.ListCompanyEventsRequest{
    Identifier: roicai.String(
        "NASDAQ:AAPL",
    ),
}
client.CompanyEvents.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**identifier:** `*string` — Restrict results to a single company. Accepts an exchange and symbol separated by a colon (e.g. `NASDAQ:AAPL`), a ticker ID, or a FIGI; the response `symbol` echoes the listing you requested. Omit to list events across all companies.
    
</dd>
</dl>

<dl>
<dd>

**content:** `*roicai.ListCompanyEventsContent` 
    
</dd>
</dl>

<dl>
<dd>

**order:** `*roicai.Order` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Maximum number of resources to return. Defaults to 50.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*string` — Opaque page token from a previous response. Prefer following the returned page URLs directly.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Calendar Earnings
<details><summary><code>client.Calendar.Earnings.List() -> *roicai.EarningsEventList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a page of dated earnings events across all tickers, or one ticker via the identifier filter. Each event carries the scheduled release date and session plus EPS and revenue estimates, actuals, and surprises.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &calendar.ListEarningsCalendarRequest{
    Identifier: roicai.String(
        "NASDAQ:AAPL",
    ),
}
client.Calendar.Earnings.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**identifier:** `*string` — Restrict results to a single ticker. Accepts an exchange and symbol separated by a colon (e.g. `NASDAQ:AAPL`), a ticker ID, or a FIGI. Omit to list earnings events across all tickers.
    
</dd>
</dl>

<dl>
<dd>

**date:** `*string` — Return events whose release date is exactly this date.
    
</dd>
</dl>

<dl>
<dd>

**dateGte:** `*string` — Return events whose release date is on or after this date.
    
</dd>
</dl>

<dl>
<dd>

**dateGt:** `*string` — Return events whose release date is strictly after this date.
    
</dd>
</dl>

<dl>
<dd>

**dateLte:** `*string` — Return events whose release date is on or before this date.
    
</dd>
</dl>

<dl>
<dd>

**dateLt:** `*string` — Return events whose release date is strictly before this date.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*roicai.Order` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Maximum number of resources to return. Defaults to 100.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*string` — Opaque page token from a previous response. Prefer following the returned page URLs directly.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Companies Profiles
<details><summary><code>client.Companies.Profiles.Retrieve(Identifier) -> *roicai.CompanyProfile</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a primary-listing company profile by public profile ID or supported market identifier.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &companies.RetrieveCompanyProfileRequest{
    Identifier: "NASDAQ:AAPL",
}
client.Companies.Profiles.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**identifier:** `string` — A company-profile ID, ticker ID, an exchange and symbol separated by a colon (e.g. `NASDAQ:AAPL`), or a FIGI.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Companies Employees
<details><summary><code>client.Companies.Employees.List() -> *roicai.EmployeeCountList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a page of fiscal-year employee counts across all companies, or one company via the identifier filter. Any listing identifier resolves to the primary company record.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &companies.ListEmployeeCountsRequest{
    Identifier: roicai.String(
        "NASDAQ:AAPL",
    ),
}
client.Companies.Employees.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**identifier:** `*string` — Restrict results to a single company. Accepts a company-profile ID, a ticker ID, an exchange and symbol separated by a colon (e.g. `NASDAQ:AAPL`), or a FIGI. Employee counts are company-level data, so secondary listings resolve to the primary listing. Omit to list counts across all companies.
    
</dd>
</dl>

<dl>
<dd>

**fiscalYear:** `*int` — Return only counts for exactly this fiscal year.
    
</dd>
</dl>

<dl>
<dd>

**fiscalYearGte:** `*int` — Return counts from this fiscal year onward.
    
</dd>
</dl>

<dl>
<dd>

**fiscalYearGt:** `*int` — Return counts after this fiscal year.
    
</dd>
</dl>

<dl>
<dd>

**fiscalYearLte:** `*int` — Return counts through this fiscal year.
    
</dd>
</dl>

<dl>
<dd>

**fiscalYearLt:** `*int` — Return counts before this fiscal year.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*roicai.Order` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Maximum number of resources to return. Defaults to 100.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*string` — Opaque page token from a previous response. Prefer following the returned page URLs directly.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Companies Logos
<details><summary><code>client.Companies.Logos.Retrieve(Identifier) -> string</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a resized company logo as WebP or PNG image bytes. Secondary listings resolve to the primary company logo.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &companies.RetrieveCompanyLogoRequest{
    Identifier: "identifier",
}
client.Companies.Logos.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**identifier:** `string` — A company-profile ID, ticker ID, an exchange and symbol separated by a colon (e.g. `NASDAQ:AAPL`), or a FIGI.
    
</dd>
</dl>

<dl>
<dd>

**size:** `*int` — Square image size in pixels. Defaults to 128.
    
</dd>
</dl>

<dl>
<dd>

**greyscale:** `*bool` — Whether to remove color from the logo. Defaults to false.
    
</dd>
</dl>

<dl>
<dd>

**format:** `*companies.RetrieveCompanyLogoFormat` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Documents Slides
<details><summary><code>client.Documents.Slides.Retrieve(ID) -> *roicai.SlidesDownload</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the metadata of an earnings-event slide deck together with a signed download URL that expires at `expires_at` (one hour by default). Professional and Enterprise plans only. Use the `doc_` id from the event’s `slides` list.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &documents.RetrieveSlidesRequest{
    ID: "doc_4JlukNobJIA1qj",
}
client.Documents.Slides.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The document id from the event listing (`doc_...`).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Documents Filings
<details><summary><code>client.Documents.Filings.Retrieve(ID) -> *roicai.DocumentDownload</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the metadata of an earnings-event filing (earnings release, quarterly or annual report, and similar) together with a signed download URL that expires at `expires_at` (one hour by default). Professional and Enterprise plans only. Use the `doc_` id from the event’s `filings` list.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &documents.RetrieveFilingRequest{
    ID: "doc_4JlukNobJIA1qj",
}
client.Documents.Filings.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The document id from the event listing (`doc_...`).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Documents Audio
<details><summary><code>client.Documents.Audio.Retrieve(ID) -> *roicai.AudioDownload</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the metadata of an earnings-call audio recording together with a signed download URL that expires at `expires_at` (one hour by default). Professional and Enterprise plans only. Use the `audio_` id from the event’s `audio` list.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &documents.RetrieveAudioRequest{
    ID: "audio_7Ft0XPWQVpVeJ1",
}
client.Documents.Audio.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The audio id from the event listing (`audio_...`).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Documents Transcripts Raw
<details><summary><code>client.Documents.Transcripts.Raw.Retrieve(ID) -> *roicai.TranscriptRaw</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the plain-text transcript of an earnings event, without speaker labels or line breaks. Free plans can read each company’s latest two events; Individual plans follow the plan’s lookback; Professional and Enterprise plans read everything. Use the `trraw_` id from the event listing.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &transcripts.RetrieveRawTranscriptRequest{
    ID: "trraw_hbBlbCUawzMoz0",
}
client.Documents.Transcripts.Raw.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The transcript id from the event listing: `trraw_...` for the raw edition, `tredit_...` for the edited edition. Both are keyed by the event and stay stable when the source changes.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Documents Transcripts Edited
<details><summary><code>client.Documents.Transcripts.Edited.Retrieve(ID) -> *roicai.TranscriptEdited</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the speaker-labelled transcript of an earnings event: ordered segments with a `speaker_id`, and a `speakers` map with each speaker’s name, role, and company where known. The human-edited edition is preferred when one exists. Professional and Enterprise plans only. Use the `tredit_` id from the event listing.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &transcripts.RetrieveEditedTranscriptRequest{
    ID: "trraw_hbBlbCUawzMoz0",
}
client.Documents.Transcripts.Edited.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` — The transcript id from the event listing: `trraw_...` for the raw edition, `tredit_...` for the edited edition. Both are keyed by the event and stay stable when the source changes.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## EarningsCalls Bulk
<details><summary><code>client.EarningsCalls.Bulk.List() -> *roicai.BulkFileList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists pre-generated transcript CSV files partitioned by fiscal year and quarter, each with a short-lived signed download URL. Filter with the optional year and period parameters. Requires a Professional or Enterprise plan.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &earningscalls.ListBulkEarningsCallsRequest{}
client.EarningsCalls.Bulk.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**year:** `*int` — Return only the files of this year. Fiscal year for fundamentals and earnings calls, calendar year for stock prices.
    
</dd>
</dl>

<dl>
<dd>

**period:** `*earningscalls.ListBulkEarningsCallsPeriod` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Exchanges TradingHours
<details><summary><code>client.Exchanges.TradingHours.Retrieve() -> *roicai.ExchangeTradingHours</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the complete regular and extended trading-hours history of one exchange as structured arrays.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &exchanges.RetrieveExchangeTradingHoursRequest{
    Exchange: "exchange",
}
client.Exchanges.TradingHours.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**exchange:** `string` — Exchange code to retrieve trading-hours history for, e.g. NASDAQ.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Exchanges Holidays
<details><summary><code>client.Exchanges.Holidays.Retrieve() -> *roicai.ExchangeHolidays</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the holiday calendar of one exchange by exchange code, as YYYY-MM-DD dates.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &exchanges.RetrieveExchangeHolidaysRequest{
    Exchange: "exchange",
}
client.Exchanges.Holidays.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**exchange:** `string` — Exchange code to retrieve the holiday calendar for, e.g. NASDAQ.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## ForexPrices Latest
<details><summary><code>client.ForexPrices.Latest.List() -> *roicai.LatestForexPriceList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a page with one row per currency pair carrying its latest available daily OHLC quote from the last 10 days. Pairs without a price in that window appear with null price fields. Requires an Individual or higher plan.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &forexprices.ListLatestForexPricesRequest{
    BaseCurrency: roicai.String(
        "EUR",
    ),
    QuoteCurrency: roicai.String(
        "EUR",
    ),
    Currency: roicai.String(
        "EUR",
    ),
}
client.ForexPrices.Latest.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**baseCurrency:** `*string` — Return only pairs quoted from this base currency.
    
</dd>
</dl>

<dl>
<dd>

**quoteCurrency:** `*string` — Return only pairs quoted into this quote currency.
    
</dd>
</dl>

<dl>
<dd>

**currency:** `*string` — Return every pair in which this currency appears, on either side.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*roicai.Order` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Maximum number of resources to return. Defaults to 500.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*string` — Opaque page token from a previous response. Prefer following the returned page URLs directly.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ForexPrices.Latest.Retrieve(Identifier) -> *roicai.ForexPrice</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the latest available daily price for a currency pair, addressed by bare currency-pair code or qualified pair symbol.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &forexprices.RetrieveLatestForexPriceRequest{
    Identifier: "EURUSD",
}
client.ForexPrices.Latest.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**identifier:** `string` — A currency pair as `BASEQUOTE` or `BASE/QUOTE`, or a qualified pair symbol (e.g. `FX:EURUSD`). The `symbol` of every `/forex-pairs` result is accepted here.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Fundamentals IncomeStatements
<details><summary><code>client.Fundamentals.IncomeStatements.List(Identifier) -> *roicai.IncomeStatementFinancialPeriodList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns income-statement line items by financial period.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &fundamentals.ListIncomeStatementsRequest{
    Identifier: "NASDAQ:AAPL",
}
client.Fundamentals.IncomeStatements.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**identifier:** `string` — A ticker ID, an exchange and symbol separated by a colon (e.g. `NASDAQ:AAPL`), or a FIGI.
    
</dd>
</dl>

<dl>
<dd>

**periodType:** `*roicai.PeriodType` 
    
</dd>
</dl>

<dl>
<dd>

**periodLabel:** `*roicai.PeriodLabel` 
    
</dd>
</dl>

<dl>
<dd>

**date:** `*string` — Return statements whose period end date is exactly this date.
    
</dd>
</dl>

<dl>
<dd>

**dateGte:** `*string` — Return statements whose period end date is on or after this date.
    
</dd>
</dl>

<dl>
<dd>

**dateGt:** `*string` — Return statements whose period end date is strictly after this date.
    
</dd>
</dl>

<dl>
<dd>

**dateLte:** `*string` — Return statements whose period end date is on or before this date.
    
</dd>
</dl>

<dl>
<dd>

**dateLt:** `*string` — Return statements whose period end date is strictly before this date.
    
</dd>
</dl>

<dl>
<dd>

**fiscalYearGte:** `*int` — Return statements from this fiscal year onward.
    
</dd>
</dl>

<dl>
<dd>

**fiscalYearGt:** `*int` — Return statements after this fiscal year.
    
</dd>
</dl>

<dl>
<dd>

**fiscalYearLte:** `*int` — Return statements through this fiscal year.
    
</dd>
</dl>

<dl>
<dd>

**fiscalYearLt:** `*int` — Return statements before this fiscal year.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*roicai.Order` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Maximum number of resources to return. Defaults to 50.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*string` — Opaque page token from a previous response. Prefer following the returned page URLs directly.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Fundamentals BalanceSheets
<details><summary><code>client.Fundamentals.BalanceSheets.List(Identifier) -> *roicai.BalanceSheetFinancialPeriodList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns balance-sheet line items by financial period.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &fundamentals.ListBalanceSheetsRequest{
    Identifier: "NASDAQ:AAPL",
}
client.Fundamentals.BalanceSheets.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**identifier:** `string` — A ticker ID, an exchange and symbol separated by a colon (e.g. `NASDAQ:AAPL`), or a FIGI.
    
</dd>
</dl>

<dl>
<dd>

**periodType:** `*roicai.PeriodType` 
    
</dd>
</dl>

<dl>
<dd>

**periodLabel:** `*roicai.PeriodLabel` 
    
</dd>
</dl>

<dl>
<dd>

**date:** `*string` — Return statements whose period end date is exactly this date.
    
</dd>
</dl>

<dl>
<dd>

**dateGte:** `*string` — Return statements whose period end date is on or after this date.
    
</dd>
</dl>

<dl>
<dd>

**dateGt:** `*string` — Return statements whose period end date is strictly after this date.
    
</dd>
</dl>

<dl>
<dd>

**dateLte:** `*string` — Return statements whose period end date is on or before this date.
    
</dd>
</dl>

<dl>
<dd>

**dateLt:** `*string` — Return statements whose period end date is strictly before this date.
    
</dd>
</dl>

<dl>
<dd>

**fiscalYearGte:** `*int` — Return statements from this fiscal year onward.
    
</dd>
</dl>

<dl>
<dd>

**fiscalYearGt:** `*int` — Return statements after this fiscal year.
    
</dd>
</dl>

<dl>
<dd>

**fiscalYearLte:** `*int` — Return statements through this fiscal year.
    
</dd>
</dl>

<dl>
<dd>

**fiscalYearLt:** `*int` — Return statements before this fiscal year.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*roicai.Order` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Maximum number of resources to return. Defaults to 50.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*string` — Opaque page token from a previous response. Prefer following the returned page URLs directly.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Fundamentals CashFlowStatements
<details><summary><code>client.Fundamentals.CashFlowStatements.List(Identifier) -> *roicai.CashFlowFinancialPeriodList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns cash-flow line items by financial period.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &fundamentals.ListCashFlowStatementsRequest{
    Identifier: "NASDAQ:AAPL",
}
client.Fundamentals.CashFlowStatements.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**identifier:** `string` — A ticker ID, an exchange and symbol separated by a colon (e.g. `NASDAQ:AAPL`), or a FIGI.
    
</dd>
</dl>

<dl>
<dd>

**periodType:** `*roicai.PeriodType` 
    
</dd>
</dl>

<dl>
<dd>

**periodLabel:** `*roicai.PeriodLabel` 
    
</dd>
</dl>

<dl>
<dd>

**date:** `*string` — Return statements whose period end date is exactly this date.
    
</dd>
</dl>

<dl>
<dd>

**dateGte:** `*string` — Return statements whose period end date is on or after this date.
    
</dd>
</dl>

<dl>
<dd>

**dateGt:** `*string` — Return statements whose period end date is strictly after this date.
    
</dd>
</dl>

<dl>
<dd>

**dateLte:** `*string` — Return statements whose period end date is on or before this date.
    
</dd>
</dl>

<dl>
<dd>

**dateLt:** `*string` — Return statements whose period end date is strictly before this date.
    
</dd>
</dl>

<dl>
<dd>

**fiscalYearGte:** `*int` — Return statements from this fiscal year onward.
    
</dd>
</dl>

<dl>
<dd>

**fiscalYearGt:** `*int` — Return statements after this fiscal year.
    
</dd>
</dl>

<dl>
<dd>

**fiscalYearLte:** `*int` — Return statements through this fiscal year.
    
</dd>
</dl>

<dl>
<dd>

**fiscalYearLt:** `*int` — Return statements before this fiscal year.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*roicai.Order` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Maximum number of resources to return. Defaults to 50.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*string` — Opaque page token from a previous response. Prefer following the returned page URLs directly.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Fundamentals EnterpriseValues
<details><summary><code>client.Fundamentals.EnterpriseValues.List(Identifier) -> *roicai.EnterpriseValueFinancialPeriodList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns enterprise values by financial period.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &fundamentals.ListEnterpriseValuesRequest{
    Identifier: "NASDAQ:AAPL",
}
client.Fundamentals.EnterpriseValues.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**identifier:** `string` — A ticker ID, an exchange and symbol separated by a colon (e.g. `NASDAQ:AAPL`), or a FIGI.
    
</dd>
</dl>

<dl>
<dd>

**periodType:** `*roicai.PeriodType` 
    
</dd>
</dl>

<dl>
<dd>

**periodLabel:** `*roicai.PeriodLabel` 
    
</dd>
</dl>

<dl>
<dd>

**date:** `*string` — Return statements whose period end date is exactly this date.
    
</dd>
</dl>

<dl>
<dd>

**dateGte:** `*string` — Return statements whose period end date is on or after this date.
    
</dd>
</dl>

<dl>
<dd>

**dateGt:** `*string` — Return statements whose period end date is strictly after this date.
    
</dd>
</dl>

<dl>
<dd>

**dateLte:** `*string` — Return statements whose period end date is on or before this date.
    
</dd>
</dl>

<dl>
<dd>

**dateLt:** `*string` — Return statements whose period end date is strictly before this date.
    
</dd>
</dl>

<dl>
<dd>

**fiscalYearGte:** `*int` — Return statements from this fiscal year onward.
    
</dd>
</dl>

<dl>
<dd>

**fiscalYearGt:** `*int` — Return statements after this fiscal year.
    
</dd>
</dl>

<dl>
<dd>

**fiscalYearLte:** `*int` — Return statements through this fiscal year.
    
</dd>
</dl>

<dl>
<dd>

**fiscalYearLt:** `*int` — Return statements before this fiscal year.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*roicai.Order` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Maximum number of resources to return. Defaults to 50.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*string` — Opaque page token from a previous response. Prefer following the returned page URLs directly.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Fundamentals Multiples
<details><summary><code>client.Fundamentals.Multiples.List(Identifier) -> *roicai.ValuationMultiplesFinancialPeriodList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns valuation multiples by financial period.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &fundamentals.ListValuationMultiplesRequest{
    Identifier: "NASDAQ:AAPL",
}
client.Fundamentals.Multiples.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**identifier:** `string` — A ticker ID, an exchange and symbol separated by a colon (e.g. `NASDAQ:AAPL`), or a FIGI.
    
</dd>
</dl>

<dl>
<dd>

**periodType:** `*roicai.PeriodType` 
    
</dd>
</dl>

<dl>
<dd>

**periodLabel:** `*roicai.PeriodLabel` 
    
</dd>
</dl>

<dl>
<dd>

**date:** `*string` — Return statements whose period end date is exactly this date.
    
</dd>
</dl>

<dl>
<dd>

**dateGte:** `*string` — Return statements whose period end date is on or after this date.
    
</dd>
</dl>

<dl>
<dd>

**dateGt:** `*string` — Return statements whose period end date is strictly after this date.
    
</dd>
</dl>

<dl>
<dd>

**dateLte:** `*string` — Return statements whose period end date is on or before this date.
    
</dd>
</dl>

<dl>
<dd>

**dateLt:** `*string` — Return statements whose period end date is strictly before this date.
    
</dd>
</dl>

<dl>
<dd>

**fiscalYearGte:** `*int` — Return statements from this fiscal year onward.
    
</dd>
</dl>

<dl>
<dd>

**fiscalYearGt:** `*int` — Return statements after this fiscal year.
    
</dd>
</dl>

<dl>
<dd>

**fiscalYearLte:** `*int` — Return statements through this fiscal year.
    
</dd>
</dl>

<dl>
<dd>

**fiscalYearLt:** `*int` — Return statements before this fiscal year.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*roicai.Order` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Maximum number of resources to return. Defaults to 50.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*string` — Opaque page token from a previous response. Prefer following the returned page URLs directly.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Fundamentals PerShareMetrics
<details><summary><code>client.Fundamentals.PerShareMetrics.List(Identifier) -> *roicai.PerShareFinancialPeriodList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns per-share metrics by financial period.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &fundamentals.ListPerShareMetricsRequest{
    Identifier: "NASDAQ:AAPL",
}
client.Fundamentals.PerShareMetrics.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**identifier:** `string` — A ticker ID, an exchange and symbol separated by a colon (e.g. `NASDAQ:AAPL`), or a FIGI.
    
</dd>
</dl>

<dl>
<dd>

**periodType:** `*roicai.PeriodType` 
    
</dd>
</dl>

<dl>
<dd>

**periodLabel:** `*roicai.PeriodLabel` 
    
</dd>
</dl>

<dl>
<dd>

**date:** `*string` — Return statements whose period end date is exactly this date.
    
</dd>
</dl>

<dl>
<dd>

**dateGte:** `*string` — Return statements whose period end date is on or after this date.
    
</dd>
</dl>

<dl>
<dd>

**dateGt:** `*string` — Return statements whose period end date is strictly after this date.
    
</dd>
</dl>

<dl>
<dd>

**dateLte:** `*string` — Return statements whose period end date is on or before this date.
    
</dd>
</dl>

<dl>
<dd>

**dateLt:** `*string` — Return statements whose period end date is strictly before this date.
    
</dd>
</dl>

<dl>
<dd>

**fiscalYearGte:** `*int` — Return statements from this fiscal year onward.
    
</dd>
</dl>

<dl>
<dd>

**fiscalYearGt:** `*int` — Return statements after this fiscal year.
    
</dd>
</dl>

<dl>
<dd>

**fiscalYearLte:** `*int` — Return statements through this fiscal year.
    
</dd>
</dl>

<dl>
<dd>

**fiscalYearLt:** `*int` — Return statements before this fiscal year.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*roicai.Order` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Maximum number of resources to return. Defaults to 50.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*string` — Opaque page token from a previous response. Prefer following the returned page URLs directly.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Fundamentals Bulk
<details><summary><code>client.Fundamentals.Bulk.List() -> *roicai.BulkFileList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists pre-generated CSV files partitioned by fiscal year and financial cadence, each with a short-lived signed download URL. Filter with the optional year and period parameters. Requires a Professional or Enterprise plan.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &fundamentals.ListBulkFundamentalsRequest{}
client.Fundamentals.Bulk.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**year:** `*int` — Return only the files of this year. Fiscal year for fundamentals and earnings calls, calendar year for stock prices.
    
</dd>
</dl>

<dl>
<dd>

**period:** `*fundamentals.ListBulkFundamentalsPeriod` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Fundamentals Ratios Profitability
<details><summary><code>client.Fundamentals.Ratios.Profitability.List(Identifier) -> *roicai.ProfitabilityFinancialPeriodList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns profitability ratios by financial period.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ratios.ListProfitabilityRatiosRequest{
    Identifier: "NASDAQ:AAPL",
}
client.Fundamentals.Ratios.Profitability.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**identifier:** `string` — A ticker ID, an exchange and symbol separated by a colon (e.g. `NASDAQ:AAPL`), or a FIGI.
    
</dd>
</dl>

<dl>
<dd>

**periodType:** `*roicai.PeriodType` 
    
</dd>
</dl>

<dl>
<dd>

**periodLabel:** `*roicai.PeriodLabel` 
    
</dd>
</dl>

<dl>
<dd>

**dateGte:** `*string` — Return statements whose period end date is on or after this date.
    
</dd>
</dl>

<dl>
<dd>

**dateGt:** `*string` — Return statements whose period end date is strictly after this date.
    
</dd>
</dl>

<dl>
<dd>

**dateLte:** `*string` — Return statements whose period end date is on or before this date.
    
</dd>
</dl>

<dl>
<dd>

**dateLt:** `*string` — Return statements whose period end date is strictly before this date.
    
</dd>
</dl>

<dl>
<dd>

**fiscalYearGte:** `*int` — Return statements from this fiscal year onward.
    
</dd>
</dl>

<dl>
<dd>

**fiscalYearGt:** `*int` — Return statements after this fiscal year.
    
</dd>
</dl>

<dl>
<dd>

**fiscalYearLte:** `*int` — Return statements through this fiscal year.
    
</dd>
</dl>

<dl>
<dd>

**fiscalYearLt:** `*int` — Return statements before this fiscal year.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*roicai.Order` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Maximum number of resources to return. Defaults to 50.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*string` — Opaque page token from a previous response. Prefer following the returned page URLs directly.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Fundamentals Ratios Credit
<details><summary><code>client.Fundamentals.Ratios.Credit.List(Identifier) -> *roicai.CreditFinancialPeriodList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns credit and debt values by financial period.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ratios.ListCreditRatiosRequest{
    Identifier: "NASDAQ:AAPL",
}
client.Fundamentals.Ratios.Credit.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**identifier:** `string` — A ticker ID, an exchange and symbol separated by a colon (e.g. `NASDAQ:AAPL`), or a FIGI.
    
</dd>
</dl>

<dl>
<dd>

**periodType:** `*roicai.PeriodType` 
    
</dd>
</dl>

<dl>
<dd>

**periodLabel:** `*roicai.PeriodLabel` 
    
</dd>
</dl>

<dl>
<dd>

**dateGte:** `*string` — Return statements whose period end date is on or after this date.
    
</dd>
</dl>

<dl>
<dd>

**dateGt:** `*string` — Return statements whose period end date is strictly after this date.
    
</dd>
</dl>

<dl>
<dd>

**dateLte:** `*string` — Return statements whose period end date is on or before this date.
    
</dd>
</dl>

<dl>
<dd>

**dateLt:** `*string` — Return statements whose period end date is strictly before this date.
    
</dd>
</dl>

<dl>
<dd>

**fiscalYearGte:** `*int` — Return statements from this fiscal year onward.
    
</dd>
</dl>

<dl>
<dd>

**fiscalYearGt:** `*int` — Return statements after this fiscal year.
    
</dd>
</dl>

<dl>
<dd>

**fiscalYearLte:** `*int` — Return statements through this fiscal year.
    
</dd>
</dl>

<dl>
<dd>

**fiscalYearLt:** `*int` — Return statements before this fiscal year.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*roicai.Order` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Maximum number of resources to return. Defaults to 50.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*string` — Opaque page token from a previous response. Prefer following the returned page URLs directly.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Fundamentals Ratios Liquidity
<details><summary><code>client.Fundamentals.Ratios.Liquidity.List(Identifier) -> *roicai.LiquidityFinancialPeriodList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns liquidity ratios by financial period.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ratios.ListLiquidityRatiosRequest{
    Identifier: "NASDAQ:AAPL",
}
client.Fundamentals.Ratios.Liquidity.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**identifier:** `string` — A ticker ID, an exchange and symbol separated by a colon (e.g. `NASDAQ:AAPL`), or a FIGI.
    
</dd>
</dl>

<dl>
<dd>

**periodType:** `*roicai.PeriodType` 
    
</dd>
</dl>

<dl>
<dd>

**periodLabel:** `*roicai.PeriodLabel` 
    
</dd>
</dl>

<dl>
<dd>

**date:** `*string` — Return statements whose period end date is exactly this date.
    
</dd>
</dl>

<dl>
<dd>

**dateGte:** `*string` — Return statements whose period end date is on or after this date.
    
</dd>
</dl>

<dl>
<dd>

**dateGt:** `*string` — Return statements whose period end date is strictly after this date.
    
</dd>
</dl>

<dl>
<dd>

**dateLte:** `*string` — Return statements whose period end date is on or before this date.
    
</dd>
</dl>

<dl>
<dd>

**dateLt:** `*string` — Return statements whose period end date is strictly before this date.
    
</dd>
</dl>

<dl>
<dd>

**fiscalYearGte:** `*int` — Return statements from this fiscal year onward.
    
</dd>
</dl>

<dl>
<dd>

**fiscalYearGt:** `*int` — Return statements after this fiscal year.
    
</dd>
</dl>

<dl>
<dd>

**fiscalYearLte:** `*int` — Return statements through this fiscal year.
    
</dd>
</dl>

<dl>
<dd>

**fiscalYearLt:** `*int` — Return statements before this fiscal year.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*roicai.Order` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Maximum number of resources to return. Defaults to 50.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*string` — Opaque page token from a previous response. Prefer following the returned page URLs directly.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Fundamentals Ratios WorkingCapital
<details><summary><code>client.Fundamentals.Ratios.WorkingCapital.List(Identifier) -> *roicai.WorkingCapitalFinancialPeriodList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns working-capital ratios by financial period.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ratios.ListWorkingCapitalRatiosRequest{
    Identifier: "NASDAQ:AAPL",
}
client.Fundamentals.Ratios.WorkingCapital.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**identifier:** `string` — A ticker ID, an exchange and symbol separated by a colon (e.g. `NASDAQ:AAPL`), or a FIGI.
    
</dd>
</dl>

<dl>
<dd>

**periodType:** `*roicai.PeriodType` 
    
</dd>
</dl>

<dl>
<dd>

**periodLabel:** `*roicai.PeriodLabel` 
    
</dd>
</dl>

<dl>
<dd>

**date:** `*string` — Return statements whose period end date is exactly this date.
    
</dd>
</dl>

<dl>
<dd>

**dateGte:** `*string` — Return statements whose period end date is on or after this date.
    
</dd>
</dl>

<dl>
<dd>

**dateGt:** `*string` — Return statements whose period end date is strictly after this date.
    
</dd>
</dl>

<dl>
<dd>

**dateLte:** `*string` — Return statements whose period end date is on or before this date.
    
</dd>
</dl>

<dl>
<dd>

**dateLt:** `*string` — Return statements whose period end date is strictly before this date.
    
</dd>
</dl>

<dl>
<dd>

**fiscalYearGte:** `*int` — Return statements from this fiscal year onward.
    
</dd>
</dl>

<dl>
<dd>

**fiscalYearGt:** `*int` — Return statements after this fiscal year.
    
</dd>
</dl>

<dl>
<dd>

**fiscalYearLte:** `*int` — Return statements through this fiscal year.
    
</dd>
</dl>

<dl>
<dd>

**fiscalYearLt:** `*int` — Return statements before this fiscal year.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*roicai.Order` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Maximum number of resources to return. Defaults to 50.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*string` — Opaque page token from a previous response. Prefer following the returned page URLs directly.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Fundamentals Ratios YieldAnalysis
<details><summary><code>client.Fundamentals.Ratios.YieldAnalysis.List(Identifier) -> *roicai.YieldAnalysisFinancialPeriodList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns yield-analysis ratios by financial period.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ratios.ListYieldAnalysisRatiosRequest{
    Identifier: "NASDAQ:AAPL",
}
client.Fundamentals.Ratios.YieldAnalysis.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**identifier:** `string` — A ticker ID, an exchange and symbol separated by a colon (e.g. `NASDAQ:AAPL`), or a FIGI.
    
</dd>
</dl>

<dl>
<dd>

**periodType:** `*roicai.PeriodType` 
    
</dd>
</dl>

<dl>
<dd>

**periodLabel:** `*roicai.PeriodLabel` 
    
</dd>
</dl>

<dl>
<dd>

**date:** `*string` — Return statements whose period end date is exactly this date.
    
</dd>
</dl>

<dl>
<dd>

**dateGte:** `*string` — Return statements whose period end date is on or after this date.
    
</dd>
</dl>

<dl>
<dd>

**dateGt:** `*string` — Return statements whose period end date is strictly after this date.
    
</dd>
</dl>

<dl>
<dd>

**dateLte:** `*string` — Return statements whose period end date is on or before this date.
    
</dd>
</dl>

<dl>
<dd>

**dateLt:** `*string` — Return statements whose period end date is strictly before this date.
    
</dd>
</dl>

<dl>
<dd>

**fiscalYearGte:** `*int` — Return statements from this fiscal year onward.
    
</dd>
</dl>

<dl>
<dd>

**fiscalYearGt:** `*int` — Return statements after this fiscal year.
    
</dd>
</dl>

<dl>
<dd>

**fiscalYearLte:** `*int` — Return statements through this fiscal year.
    
</dd>
</dl>

<dl>
<dd>

**fiscalYearLt:** `*int` — Return statements before this fiscal year.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*roicai.Order` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Maximum number of resources to return. Defaults to 50.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*string` — Opaque page token from a previous response. Prefer following the returned page URLs directly.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## StockPrices Latest
<details><summary><code>client.StockPrices.Latest.List() -> *roicai.LatestStockPriceList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a page with one row per ticker carrying its latest available daily OHLCV quote from the last 10 days. Tickers without a price in that window (halted, delisted, or unpriced listings) appear with null price fields. Requires an Individual or higher plan.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &stockprices.ListLatestStockPricesRequest{}
client.StockPrices.Latest.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**exchange:** `*string` — Filter to tickers listed on one exchange code, e.g. `NASDAQ`.
    
</dd>
</dl>

<dl>
<dd>

**order:** `*roicai.Order` 
    
</dd>
</dl>

<dl>
<dd>

**limit:** `*int` — Maximum number of resources to return. Defaults to 500.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*string` — Opaque page token from a previous response. Prefer following the returned page URLs directly.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.StockPrices.Latest.Retrieve(Identifier) -> *roicai.HistoricalPrice</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the latest available daily price for a ticker, symbol, or other supported security identifier.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &stockprices.RetrieveLatestStockPriceRequest{
    Identifier: "NASDAQ:AAPL",
}
client.StockPrices.Latest.Retrieve(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**identifier:** `string` — A ticker ID, an exchange and symbol separated by a colon (e.g. `NASDAQ:AAPL`), or a FIGI.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## StockPrices Bulk
<details><summary><code>client.StockPrices.Bulk.List() -> *roicai.BulkFileList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Lists pre-generated daily-price CSV files partitioned by calendar year and month, each with a short-lived signed download URL. Filter with the optional year and period parameters. Requires a Professional or Enterprise plan.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &stockprices.ListBulkStockPricesRequest{}
client.StockPrices.Bulk.List(
    context.TODO(),
    request,
)
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**year:** `*int` — Return only the files of this year. Fiscal year for fundamentals and earnings calls, calendar year for stock prices.
    
</dd>
</dl>

<dl>
<dd>

**period:** `*stockprices.ListBulkStockPricesPeriod` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

