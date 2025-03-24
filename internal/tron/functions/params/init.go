package params

import (
	"errors"
	"github.com/wenpiner/tron-scan/internal/types"
	"strconv"
)

type FunctionParam struct {
	Type       string
	Len        int
	dataHandle func(data string) (interface{}, error)
}

func (param *FunctionParam) GetValue(data string) (interface{}, error) {
	// 截取字符串
	if len(data) < param.Len {
		return nil, errors.New("data length is not enough")
	}
	hex := data[:param.Len]
	return param.dataHandle(hex)
}

var FunctionParams = map[string]FunctionParam{
	"address": {
		Type: "address",
		Len:  64,
		dataHandle: func(data string) (interface{}, error) {
			return types.AddressByHex(data).String(), nil
		},
	},
	"uint256": {
		Type: "uint256",
		Len:  64,
		dataHandle: func(data string) (interface{}, error) {
			return strconv.ParseInt(data, 16, 64)
		},
	},
}
