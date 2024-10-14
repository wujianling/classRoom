package classroomTalk

import (
	"context"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelClassRoomTalkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelClassRoomTalkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelClassRoomTalkLogic {
	return &DelClassRoomTalkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelClassRoomTalkLogic) DelClassRoomTalk(req *types.DelClassRoomTalkReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
