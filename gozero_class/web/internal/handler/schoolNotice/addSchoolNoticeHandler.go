package schoolNotice

import (
	"net/http"

	"github.com/smallq_class/web/internal/logic/schoolNotice"
	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddSchoolNoticeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddSchoolNoticeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := schoolNotice.NewAddSchoolNoticeLogic(r.Context(), svcCtx)
		resp, err := l.AddSchoolNotice(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
