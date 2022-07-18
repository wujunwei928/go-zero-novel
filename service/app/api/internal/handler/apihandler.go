package handler

import (
	"net/http"

	"github.com/wujunwei928/go-zero-novel/service/app/api/internal/logic"
	"github.com/wujunwei928/go-zero-novel/service/app/api/internal/svc"
	"github.com/wujunwei928/go-zero-novel/service/app/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ApiHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewApiLogic(r.Context(), svcCtx)
		resp, err := l.Api(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
