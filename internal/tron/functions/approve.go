package functions

import (
	"github.com/wenpiner/tron-scan/internal/tron/functions/params"
	"github.com/wenpiner/tron-scan/internal/types"
)

type ApproveFunction struct {
	params []params.FunctionParam
}

func (a *ApproveFunction) HandleMessage(id string, num uint64, hash string, timestamp int64, c *types.TriggerSmartContract) (*types.MQTransactionMessage, error) {
	return &types.MQTransactionMessage{
		Hash:         id,
		BlockNum:     num,
		Amount:       c.DataInfo.Values[1].(int64),
		Contract:     c.ContractAddress.String(),
		BlockHash:    hash,
		FromAddr:     c.OwnerAddress.String(),
		ToAddr:       c.DataInfo.Values[0].(string),
		Datetime:     timestamp,
		FunctionName: "approve",
		Values:       c.DataInfo.Values,
	}, nil
}

func (a *ApproveFunction) Params() []params.FunctionParam {
	return a.params
}

func (a *ApproveFunction) FunctionName() string {
	return "approve"
}

func init() {
	RegisterFunction(&ApproveFunction{
		params: []params.FunctionParam{
			params.FunctionParams["address"],
			params.FunctionParams["uint256"],
		},
	})
}
