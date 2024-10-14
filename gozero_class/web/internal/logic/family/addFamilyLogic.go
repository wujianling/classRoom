package family

import (
	"context"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddFamilyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddFamilyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFamilyLogic {
	return &AddFamilyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddFamilyLogic) AddFamily(req *types.AddFamilyReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
