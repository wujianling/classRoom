package user

import (
	"net/http"

	"github.com/smallq_class/app/internal/logic/user"
	"github.com/smallq_class/app/internal/svc"
	"github.com/smallq_class/app/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserGetSMSCodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserGetSMSCodeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewUserGetSMSCodeLogic(r.Context(), svcCtx)
		resp, err := l.UserGetSMSCode(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
