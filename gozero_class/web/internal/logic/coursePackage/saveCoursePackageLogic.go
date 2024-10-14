package coursePackage

import (
	"context"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveCoursePackageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveCoursePackageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveCoursePackageLogic {
	return &SaveCoursePackageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveCoursePackageLogic) SaveCoursePackage(req *types.SaveCoursePackageReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
