package courseLog

import (
	"context"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchCourseLogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchCourseLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchCourseLogListLogic {
	return &SearchCourseLogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchCourseLogListLogic) SearchCourseLogList(req *types.SearchCourseLogListReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
