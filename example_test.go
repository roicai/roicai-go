// Hand-written: examples shown on pkg.go.dev. They compile in `go test` but
// don't run, since they call the live API.

package roicai_test

import (
	"context"
	"errors"
	"fmt"
	"log"

	roicai "github.com/roicai/roicai-go"
	"github.com/roicai/roicai-go/client"
	"github.com/roicai/roicai-go/fundamentals"
	"github.com/roicai/roicai-go/option"
)

func Example() {
	// Reads the API key from ROICAI_API_KEY.
	c := client.NewClient()
	ctx := context.Background()

	ticker, err := c.Tickers.Retrieve(ctx, &roicai.RetrieveTickerRequest{Identifier: "NASDAQ:AAPL"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*ticker.Name)

	statements, err := c.Fundamentals.IncomeStatements.List(ctx, &fundamentals.ListIncomeStatementsRequest{
		Identifier: "NASDAQ:AAPL",
		PeriodType: roicai.PeriodTypeAnnual.Ptr(),
		Limit:      roicai.Int(5),
	})
	if err != nil {
		log.Fatal(err)
	}
	for _, year := range statements.Results {
		fmt.Println(*year.FiscalYear, *year.IsSalesRevenueTurnover)
	}
}

func Example_withAPIKey() {
	c := client.NewClient(option.WithAPIKey("your-api-key"))

	_, _ = c.Tickers.Retrieve(context.Background(), &roicai.RetrieveTickerRequest{Identifier: "NASDAQ:AAPL"})
}

func Example_pagination() {
	c := client.NewClient()
	ctx := context.Background()

	page, err := c.StockPrices.List(ctx, &roicai.ListHistoricalStockPricesRequest{
		Identifier: "NASDAQ:AAPL",
		DateGte:    roicai.String("2020-01-01"),
	})
	if err != nil {
		log.Fatal(err)
	}
	// All fetches each next page when the loop reaches it.
	for price, err := range page.All(ctx) {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(price.Date.Format("2006-01-02"), price.Close)
	}
}

func Example_errors() {
	c := client.NewClient()

	_, err := c.Tickers.Retrieve(context.Background(), &roicai.RetrieveTickerRequest{Identifier: "NASDAQ:NOPE"})

	var notFound *roicai.NotFoundError
	if errors.As(err, &notFound) {
		fmt.Println(notFound.Body.GetError().GetCode()) // resource_missing
	}
}
