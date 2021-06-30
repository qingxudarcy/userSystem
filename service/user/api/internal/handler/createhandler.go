package handler

import (
	"net/http"

	"userSystem/service/user/api/internal/logic"
	"userSystem/service/user/api/internal/svc"
	"userSystem/service/user/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func createHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCreateLogic(r.Context(), ctx)
		resp, err := l.Create(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
