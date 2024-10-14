package classroomStu

import (
	"context"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveClassRoomStuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveClassRoomStuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveClassRoomStuLogic {
	return &SaveClassRoomStuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveClassRoomStuLogic) SaveClassRoomStu(req *types.SaveClassRoomStuReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
