package parent

import (
	"context"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveParentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveParentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveParentLogic {
	return &SaveParentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveParentLogic) SaveParent(req *types.SaveParentReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
