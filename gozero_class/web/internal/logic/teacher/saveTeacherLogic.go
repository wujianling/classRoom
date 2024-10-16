package teacher

import (
	"context"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveTeacherLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveTeacherLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveTeacherLogic {
	return &SaveTeacherLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveTeacherLogic) SaveTeacher(req *types.SaveTeacherReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
