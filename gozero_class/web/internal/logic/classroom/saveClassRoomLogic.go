package classroom

import (
	"context"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveClassRoomLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveClassRoomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveClassRoomLogic {
	return &SaveClassRoomLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveClassRoomLogic) SaveClassRoom(req *types.SaveClassRoomReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
