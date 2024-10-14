package family

import (
	"context"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelFamilyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelFamilyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelFamilyLogic {
	return &DelFamilyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelFamilyLogic) DelFamily(req *types.DelFamilyReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
