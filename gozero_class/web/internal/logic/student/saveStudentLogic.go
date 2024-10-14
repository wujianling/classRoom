package student

import (
	"context"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveStudentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveStudentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveStudentLogic {
	return &SaveStudentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveStudentLogic) SaveStudent(req *types.SaveStudentReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
