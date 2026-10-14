// Hand-written.

// Package roicai is the official Go client for the ROIC.ai financial data and
// stock API: fundamentals such as income statements, balance sheets, cash
// flow statements, ratios and valuation multiples, earnings call transcripts,
// stock and forex prices, SEC filings and company profiles.
//
// This package holds the request and response types. Create a client with
// [github.com/roicai/roicai-go/client.NewClient], which reads your API key
// from ROICAI_API_KEY:
//
//	c := client.NewClient()
//	ticker, err := c.Tickers.Retrieve(ctx, &roicai.RetrieveTickerRequest{Identifier: "NASDAQ:AAPL"})
//
// Methods are grouped the way the API's URLs are, so
// /v3.0.0/fundamental/ratios/credit/{identifier} is
// c.Fundamentals.Ratios.Credit.List. See https://api.roic.ai/v3.0.0/openapi
// for every endpoint, parameter and field.
package roicai
