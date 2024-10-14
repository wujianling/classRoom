package coursePackage

import (
	"context"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelCoursePackageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelCoursePackageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelCoursePackageLogic {
	return &DelCoursePackageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelCoursePackageLogic) DelCoursePackage(req *types.DelCoursePackageReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
