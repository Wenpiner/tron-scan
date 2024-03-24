package tron

import (
	"errors"
	"github.com/go-resty/resty/v2"
	"testing"
	"tronScan/internal/types"
)

func getBlock() (*types.Block, error) {
	request := map[string]interface{}{
		"visible": true,
	}
	url := "https://api.trongrid.io/wallet/getnowblock"
	var block types.Block
	response, err := resty.New().R().SetResult(&block).SetBody(request).Post(url)
	if err != nil {
		return nil, err
	}
	if response.IsSuccess() {
		return &block, nil
	} else {
		return nil, errors.New(response.String())
	}
}

func TestGetBlock(t *testing.T) {
	block, err := getBlock()
	if err != nil {
		t.Error(err)
	}
	t.Log(block)
}
