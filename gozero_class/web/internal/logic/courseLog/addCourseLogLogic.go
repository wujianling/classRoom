package courseLog

import (
	"context"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCourseLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddCourseLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCourseLogLogic {
	return &AddCourseLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCourseLogLogic) AddCourseLog(req *types.AddCourseLogReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
