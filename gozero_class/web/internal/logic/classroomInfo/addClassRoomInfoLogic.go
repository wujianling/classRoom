package classroomInfo

import (
	"context"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddClassRoomInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddClassRoomInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddClassRoomInfoLogic {
	return &AddClassRoomInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddClassRoomInfoLogic) AddClassRoomInfo(req *types.AddClassRoomInfoReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
