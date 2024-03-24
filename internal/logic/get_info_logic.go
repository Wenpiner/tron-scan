package logic

import (
	"context"
	"time"

	"tronScan/internal/svc"
	"tronScan/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetInfoLogic {
	return &GetInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetInfoLogic) GetInfo() (resp *types.InfoResponse, err error) {
	l.svcCtx.Info.CurrentTime = time.Now().UTC().Unix()
	resp = &l.svcCtx.Info
	return
}
