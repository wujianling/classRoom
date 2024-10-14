package courseLog

import (
	"context"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchCourseLogInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchCourseLogInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchCourseLogInfoLogic {
	return &SearchCourseLogInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchCourseLogInfoLogic) SearchCourseLogInfo(req *types.SearchCourseLogInfoReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
