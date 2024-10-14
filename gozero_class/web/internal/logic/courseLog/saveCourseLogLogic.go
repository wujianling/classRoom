package courseLog

import (
	"context"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveCourseLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveCourseLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveCourseLogLogic {
	return &SaveCourseLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveCourseLogLogic) SaveCourseLog(req *types.SaveCourseLogReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
