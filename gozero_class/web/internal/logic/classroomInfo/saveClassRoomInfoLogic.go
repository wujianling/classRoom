package classroomInfo

import (
	"context"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveClassRoomInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveClassRoomInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveClassRoomInfoLogic {
	return &SaveClassRoomInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveClassRoomInfoLogic) SaveClassRoomInfo(req *types.SaveClassRoomInfoReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
