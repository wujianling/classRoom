package classroom

import (
	"context"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddClassRoomLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddClassRoomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddClassRoomLogic {
	return &AddClassRoomLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddClassRoomLogic) AddClassRoom(req *types.AddClassRoomReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
