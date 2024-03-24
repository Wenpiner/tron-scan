package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"time"

	"tronScan/internal/svc"
	"tronScan/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HealthCheckLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHealthCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HealthCheckLogic {
	return &HealthCheckLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HealthCheckLogic) HealthCheck(req *types.HealthRequest) (resp *types.InfoResponse, err error) {
	now := time.Now().UTC().Unix() - l.svcCtx.Info.BlockLastUpdateTime
	if now >= req.Timeout {
		// 超时
		err = status.Error(500, "timeout")
	} else {
		l.svcCtx.Info.CurrentTime = time.Now().UTC().Unix()
		resp = &l.svcCtx.Info
	}
	return
}
