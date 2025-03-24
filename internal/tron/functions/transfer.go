package functions

import (
	"github.com/wenpiner/tron-scan/internal/tron/functions/params"
	"github.com/wenpiner/tron-scan/internal/types"
)

type TransferFunction struct {
	params []params.FunctionParam
}

func (f *TransferFunction) HandleMessage(id string, num uint64, hash string, timestamp int64, c *types.TriggerSmartContract) (*types.MQTransactionMessage, error) {
	return &types.MQTransactionMessage{
		Hash:         id,
		BlockNum:     num,
		Amount:       c.DataInfo.Values[1].(int64),
		Contract:     c.ContractAddress.String(),
		BlockHash:    hash,
		FromAddr:     c.OwnerAddress.String(),
		ToAddr:       c.DataInfo.Values[0].(string),
		Datetime:     timestamp,
		FunctionName: "transfer",
		Values:       c.DataInfo.Values,
	}, nil
}

func (f *TransferFunction) Params() []params.FunctionParam {
	return f.params
}

func (f *TransferFunction) FunctionName() string {
	return "transfer"
}

func init() {
	// 计算函数签名
	RegisterFunction(&TransferFunction{
		params: []params.FunctionParam{
			params.FunctionParams["address"],
			params.FunctionParams["uint256"],
		},
	})
}
