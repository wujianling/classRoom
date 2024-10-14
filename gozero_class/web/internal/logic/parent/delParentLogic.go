package parent

import (
	"context"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelParentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelParentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelParentLogic {
	return &DelParentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelParentLogic) DelParent(req *types.DelParentReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
