package classroomTalk

import (
	"context"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddClassRoomTalkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddClassRoomTalkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddClassRoomTalkLogic {
	return &AddClassRoomTalkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddClassRoomTalkLogic) AddClassRoomTalk(req *types.AddClassRoomTalkReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
