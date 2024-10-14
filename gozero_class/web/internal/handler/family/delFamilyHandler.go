package family

import (
	"net/http"

	"github.com/smallq_class/web/internal/logic/family"
	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DelFamilyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DelFamilyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := family.NewDelFamilyLogic(r.Context(), svcCtx)
		resp, err := l.DelFamily(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
