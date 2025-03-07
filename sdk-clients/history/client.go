package history

import (
	"github.com/paraleipsis/1inch-sdk-go/common"
)

type Client struct {
	api
}

type api struct {
	httpExecutor common.HttpExecutor
}

func NewClient(cfg *Configuration) (*Client, error) {
	c := Client{
		api: cfg.API,
	}
	return &c, nil
}
