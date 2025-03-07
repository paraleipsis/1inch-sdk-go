package gasprices

import (
	"github.com/paraleipsis/1inch-sdk-go/common"
)

type Client struct {
	api
}

type api struct {
	chainId      uint64
	httpExecutor common.HttpExecutor
}

func NewClient(cfg *Configuration) (*Client, error) {
	c := Client{
		api: cfg.API,
	}
	return &c, nil
}
