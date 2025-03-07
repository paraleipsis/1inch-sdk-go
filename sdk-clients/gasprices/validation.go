package gasprices

import "github.com/paraleipsis/1inch-sdk-go/constants"

func isEIP1559Applicable(c uint64) bool {
	return !(c == constants.BscChainId || c == constants.AuroraChainId || c == constants.ZkSyncEraChainId || c == constants.FantomChainId)
}
