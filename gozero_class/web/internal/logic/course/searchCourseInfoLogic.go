package course

import (
	"context"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchCourseInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchCourseInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchCourseInfoLogic {
	return &SearchCourseInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchCourseInfoLogic) SearchCourseInfo(req *types.SearchCourseInfoReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
