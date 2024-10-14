package classroom

import (
	"context"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelClassRoomLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelClassRoomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelClassRoomLogic {
	return &DelClassRoomLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelClassRoomLogic) DelClassRoom(req *types.DelClassRoomReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
