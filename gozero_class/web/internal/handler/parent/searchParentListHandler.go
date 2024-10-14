package parent

import (
	"net/http"

	"github.com/smallq_class/web/internal/logic/parent"
	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SearchParentListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SearchParentListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := parent.NewSearchParentListLogic(r.Context(), svcCtx)
		resp, err := l.SearchParentList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
