package teacher

import (
	"context"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTeacherLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddTeacherLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTeacherLogic {
	return &AddTeacherLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddTeacherLogic) AddTeacher(req *types.AddTeacherReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
