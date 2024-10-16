package teacher

import (
	"context"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelTeacherLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelTeacherLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelTeacherLogic {
	return &DelTeacherLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelTeacherLogic) DelTeacher(req *types.DelTeacherReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
