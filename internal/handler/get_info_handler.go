package handler

import (
	"net/http"

	"github.com/wenpiner/tron-scan/internal/logic"
	"github.com/wenpiner/tron-scan/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetInfo()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
