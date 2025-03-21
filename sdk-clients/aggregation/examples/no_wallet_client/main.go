package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/paraleipsis/1inch-sdk-go/constants"
	"github.com/paraleipsis/1inch-sdk-go/sdk-clients/aggregation"
)

var (
	devPortalToken = os.Getenv("DEV_PORTAL_TOKEN")
)

const (
	PolygonDai  = "0x8f3Cf7ad23Cd3CaDbD9735AFf958023239c6A063"
	PolygonUsdc = "0x3c499c542cEF5E3811e1192ce70d8cC03d5c3359"
)

func main() {
	config, err := aggregation.NewConfigurationAPI(constants.PolygonChainId, "https://api.1inch.dev", devPortalToken)
	if err != nil {
		log.Fatalf("Failed to create configuration: %v\n", err)
	}
	client, err := aggregation.NewClientOnlyAPI(config)
	if err != nil {
		log.Fatalf("Failed to create client: %v\n", err)
	}
	ctx := context.Background()

	quote, err := client.GetQuote(ctx, aggregation.GetQuoteParams{
		Src:    PolygonDai,
		Dst:    PolygonUsdc,
		Amount: "1000000000000000000",
	})

	if err != nil {
		log.Fatalf("Failed to get quote: %v\n", err)
	}

	fmt.Printf("Quote Amount: %+v\n", quote.DstAmount)
}
