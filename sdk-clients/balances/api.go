package balances

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/paraleipsis/1inch-sdk-go/common"
)

// GetBalancesAndAllowancesByWalletAddressList Get balances and allowances of tokens by spender for walletAddress
func (api *api) GetBalancesAndAllowancesByWalletAddressList(ctx context.Context, params BalancesAndAllowancesByWalletAddressListParams) (*BalancesAndAllowancesByWalletAddressListResponse, error) {
	u := fmt.Sprintf("/balance/v1.2/%d/allowancesAndBalances/%s/%s", api.chainId, params.Spender, params.Wallet)

	err := params.Validate()
	if err != nil {
		return nil, err
	}

	payload := common.RequestPayload{
		Method: "GET",
		Params: nil,
		U:      u,
		Body:   nil,
	}

	var response BalancesAndAllowancesByWalletAddressListResponse
	err = api.httpExecutor.ExecuteRequest(ctx, payload, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetBalancesAndAllowances Get balances and allowances by spender for list of EVM addresses
func (api *api) GetBalancesAndAllowances(ctx context.Context, params BalancesAndAllowancesParams) (*AggregatedBalancesAndAllowancesResponse, error) {
	u := fmt.Sprintf("/balance/v1.2/%d/aggregatedBalancesAndAllowances/%s", api.chainId, params.Spender)

	err := params.Validate()
	if err != nil {
		return nil, err
	}

	payload := common.RequestPayload{
		Method: "GET",
		Params: struct {
			Wallets string `url:"wallets"`
		}{Wallets: strings.Join(params.Wallets, ",")},
		U:    u,
		Body: nil,
	}

	var response AggregatedBalancesAndAllowancesResponse
	err = api.httpExecutor.ExecuteRequest(ctx, payload, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetBalancesByWalletAddress Get balances of tokens for walletAddress for examples token list (1inch tokens list)
func (api *api) GetBalancesByWalletAddress(ctx context.Context, params BalancesByWalletAddressParams) (*BalancesByWalletAddressResponse, error) {
	u := fmt.Sprintf("/balance/v1.2/%d/balances/%s", api.chainId, params.Wallet)

	err := params.Validate()
	if err != nil {
		return nil, err
	}

	payload := common.RequestPayload{
		Method: "GET",
		Params: nil,
		U:      u,
		Body:   nil,
	}

	var response BalancesByWalletAddressResponse
	err = api.httpExecutor.ExecuteRequest(ctx, payload, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetBalancesOfCustomTokensByWalletAddress Get balances of custom tokens for walletAddress
// Takes wallet address and provided tokens and provides balances of each token
func (api *api) GetBalancesOfCustomTokensByWalletAddress(ctx context.Context, params BalancesOfCustomTokensByWalletAddressParams) (*BalancesOfCustomTokensByWalletAddressResponse, error) {
	u := fmt.Sprintf("/balance/v1.2/%d/balances/%s", api.chainId, params.Wallet)

	err := params.Validate()
	if err != nil {
		return nil, err
	}

	body, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	payload := common.RequestPayload{
		Method: "POST",
		Params: nil,
		U:      u,
		Body:   body,
	}

	var response BalancesOfCustomTokensByWalletAddressResponse
	err = api.httpExecutor.ExecuteRequest(ctx, payload, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetBalancesOfCustomTokensByWalletAddress Get balances of custom tokens for list of wallets addresses
func (api *api) GetBalancesOfCustomTokensByWalletAddressesList(ctx context.Context, params BalancesOfCustomTokensByWalletAddressesListParams) (*BalancesOfCustomTokensByWalletAddressesListResponse, error) {
	u := fmt.Sprintf("/balance/v1.2/%d/balances/multiple/walletsAndTokens", api.chainId)

	err := params.Validate()
	if err != nil {
		return nil, err
	}

	body, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	payload := common.RequestPayload{
		Method: "POST",
		Params: nil,
		U:      u,
		Body:   body,
	}

	var response BalancesOfCustomTokensByWalletAddressesListResponse
	err = api.httpExecutor.ExecuteRequest(ctx, payload, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetBalancesAndAllowancesOfCustomTokensByWalletAddressList Get balances and allowances of custom tokens by spender for walletAddress
func (api *api) GetBalancesAndAllowancesOfCustomTokensByWalletAddressList(ctx context.Context, params BalancesAndAllowancesOfCustomTokensByWalletAddressParams) (*BalancesAndAllowancesOfCustomTokensByWalletAddressResponse, error) {
	u := fmt.Sprintf("/balance/v1.2/%d/allowancesAndBalances/%s/%s", api.chainId, params.Spender, params.Wallet)

	err := params.Validate()
	if err != nil {
		return nil, err
	}

	body, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	payload := common.RequestPayload{
		Method: "POST",
		Params: nil,
		U:      u,
		Body:   body,
	}

	var response BalancesAndAllowancesOfCustomTokensByWalletAddressResponse
	err = api.httpExecutor.ExecuteRequest(ctx, payload, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetBalancesByWalletAddress Get allowances of tokens by spender for walletAddress
func (api *api) GetAllowancesByWalletAddress(ctx context.Context, params AllowancesByWalletAddressParams) (*AllowancesByWalletAddressResponse, error) {
	u := fmt.Sprintf("/balance/v1.2/%d/allowances/%s/%s", api.chainId, params.Spender, params.Wallet)

	err := params.Validate()
	if err != nil {
		return nil, err
	}

	payload := common.RequestPayload{
		Method: "GET",
		Params: nil,
		U:      u,
		Body:   nil,
	}

	var response AllowancesByWalletAddressResponse
	err = api.httpExecutor.ExecuteRequest(ctx, payload, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetAllowancesOfCustomTokensByWalletAddress Get allowances of custom tokens by spender for walletAddress
func (api *api) GetAllowancesOfCustomTokensByWalletAddress(ctx context.Context, params AllowancesOfCustomTokensByWalletAddressParams) (*AllowancesOfCustomTokensByWalletAddressResponse, error) {
	u := fmt.Sprintf("/balance/v1.2/%d/allowances/%s/%s", api.chainId, params.Spender, params.Wallet)

	err := params.Validate()
	if err != nil {
		return nil, err
	}

	body, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	payload := common.RequestPayload{
		Method: "POST",
		Params: nil,
		U:      u,
		Body:   body,
	}

	var response AllowancesOfCustomTokensByWalletAddressResponse
	err = api.httpExecutor.ExecuteRequest(ctx, payload, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
