package svc

import (
	"github.com/wenpiner/tron-scan/internal/config"
	"github.com/wenpiner/tron-scan/internal/types"
	"time"
)

type ServiceContext struct {
	Config config.Config
	Info   types.InfoResponse
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Info: types.InfoResponse{
			BlockLastUpdateTime: 0,
			BlockNum:            0,
			CurrentTime:         time.Now().UTC().Unix(),
		},
	}
}

func (ctx *ServiceContext) UpdateBlockInfo(block types.Block) {
	ctx.Info.BlockLastUpdateTime = time.Now().UTC().Unix()
	ctx.Info.BlockNum = int64(block.BlockHeader.RawData.Number)
}
