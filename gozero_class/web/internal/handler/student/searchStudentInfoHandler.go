package student

import (
	"net/http"

	"github.com/smallq_class/web/internal/logic/student"
	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SearchStudentInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SearchStudentInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := student.NewSearchStudentInfoLogic(r.Context(), svcCtx)
		resp, err := l.SearchStudentInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
