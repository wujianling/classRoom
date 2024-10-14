package course

import (
	"context"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelCourseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelCourseLogic {
	return &DelCourseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelCourseLogic) DelCourse(req *types.DelCourseReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
