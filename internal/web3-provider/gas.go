package web3_provider

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/paraleipsis/1inch-sdk-go/constants"
)

func (w Wallet) GetGasTipCap(ctx context.Context) (*big.Int, error) {
	if !w.IsEIP1559Applicable() {
		return nil, fmt.Errorf("eip1159 is not supported, use gas price")
	}
	return w.ethClient.SuggestGasTipCap(ctx)
}

func (w Wallet) GetGasPrice(ctx context.Context) (*big.Int, error) {
	return w.ethClient.SuggestGasPrice(ctx)
}

func (w Wallet) GetGasEstimate(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	return w.ethClient.EstimateGas(ctx, msg)
}

func (w Wallet) IsEIP1559Applicable() bool {
	c := w.ChainId()
	return !(c == constants.BscChainId || c == constants.AuroraChainId || c == constants.ZkSyncEraChainId || c == constants.FantomChainId)
}
