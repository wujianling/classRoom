package classroomStu

import (
	"context"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddClassRoomStuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddClassRoomStuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddClassRoomStuLogic {
	return &AddClassRoomStuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddClassRoomStuLogic) AddClassRoomStu(req *types.AddClassRoomStuReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
