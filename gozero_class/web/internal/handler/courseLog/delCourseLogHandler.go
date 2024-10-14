package courseLog

import (
	"net/http"

	"github.com/smallq_class/web/internal/logic/courseLog"
	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DelCourseLogHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DelCourseLogReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := courseLog.NewDelCourseLogLogic(r.Context(), svcCtx)
		resp, err := l.DelCourseLog(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
