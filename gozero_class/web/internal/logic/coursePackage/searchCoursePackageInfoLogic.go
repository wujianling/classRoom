package coursePackage

import (
	"context"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchCoursePackageInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchCoursePackageInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchCoursePackageInfoLogic {
	return &SearchCoursePackageInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchCoursePackageInfoLogic) SearchCoursePackageInfo(req *types.SearchCoursePackageInfoReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
