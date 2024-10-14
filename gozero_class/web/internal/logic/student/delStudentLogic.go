package student

import (
	"context"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelStudentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelStudentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelStudentLogic {
	return &DelStudentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelStudentLogic) DelStudent(req *types.DelStudentReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
