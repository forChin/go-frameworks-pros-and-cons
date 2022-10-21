package handler

import (
	"net/http"

	"github.com/forChin/test-crud/go-zero-crud/internal/logic"
	"github.com/forChin/test-crud/go-zero-crud/internal/svc"
	"github.com/forChin/test-crud/go-zero-crud/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserLogic(r.Context(), svcCtx)
		resp, err := l.User(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
