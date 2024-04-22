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

// 除去bytes前面的0值
func removeZeroBytes(bytes []byte) []byte {
	for i, b := range bytes {
		if b != 0 {
			return bytes[i:]
		}
	}
	return nil
}

func TestAddr(t *testing.T) {
	// a9059cbb0000000000000000000000418d9623f261120a5155926be91eb83352e056cf4000000000000000000000000000000000000000000000000000000002a64563e0
	data := "a9059cbb0000000000000000000000418d9623f261120a5155926be91eb83352e056cf4000000000000000000000000000000000000000000000000000000002a64563e0"
	// 截取前8位
	method := data[0:8]
	if method == "a9059cbb" {
		addr := data[8:72]
		hex := types.AddressByHex(addr)
		t.Log(hex.String())
	}
}
