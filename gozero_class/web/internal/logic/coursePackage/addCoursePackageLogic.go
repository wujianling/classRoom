package coursePackage

import (
	"context"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCoursePackageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddCoursePackageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCoursePackageLogic {
	return &AddCoursePackageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCoursePackageLogic) AddCoursePackage(req *types.AddCoursePackageReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
