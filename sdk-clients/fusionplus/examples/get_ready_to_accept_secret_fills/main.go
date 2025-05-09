package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/paraleipsis/1inch-sdk-go/sdk-clients/fusionplus"
)

var (
	devPortalToken = os.Getenv("DEV_PORTAL_TOKEN")
	publicAddress  = os.Getenv("WALLET_ADDRESS")
	privateKey     = os.Getenv("WALLET_KEY")
)

const (
	usdc    = "0x3c499c542cEF5E3811e1192ce70d8cC03d5c3359"
	wmatic  = "0x0d500B1d8E8eF31E21C99d1Db9A6444d3ADf1270"
	amount  = "100000000"
	chainId = 137
)

func main() {
	config, err := fusionplus.NewConfiguration(fusionplus.ConfigurationParams{
		ApiUrl:     "https://api.1inch.dev",
		ApiKey:     devPortalToken,
		PrivateKey: privateKey,
	})
	if err != nil {
		log.Fatalf("failed to create configuration: %v", err)
	}
	client, err := fusionplus.NewClient(config)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	ctx := context.Background()

	response, err := client.GetOrderByOrderHash(ctx, fusionplus.GetOrderByOrderHashParams{
		Hash: "0x97729858044d3838c82f2ea5ca4764bd20bfdf1f99d3af05786e4a358b16fa91",
	})
	if err != nil {
		log.Fatalf("failed to request: %v", err)
	}

	output, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal response: %v\n", err)
	}
	fmt.Printf("Response: %s\n", string(output))
}
