package classroomInfo

import (
	"context"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelClassRoomInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelClassRoomInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelClassRoomInfoLogic {
	return &DelClassRoomInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelClassRoomInfoLogic) DelClassRoomInfo(req *types.DelClassRoomInfoReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
