package parent

import (
	"context"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddParentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddParentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddParentLogic {
	return &AddParentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddParentLogic) AddParent(req *types.AddParentReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
