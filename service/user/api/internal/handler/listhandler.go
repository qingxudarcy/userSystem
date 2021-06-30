package handler

import (
	"net/http"

	"userSystem/service/user/api/internal/logic"
	"userSystem/service/user/api/internal/svc"
	"userSystem/service/user/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func listHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewListLogic(r.Context(), ctx)
		resp, err := l.List(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
