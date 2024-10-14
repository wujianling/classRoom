package classroomStu

import (
	"context"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelClassRoomStuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelClassRoomStuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelClassRoomStuLogic {
	return &DelClassRoomStuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelClassRoomStuLogic) DelClassRoomStu(req *types.DelClassRoomStuReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
