package classroomInfo

import (
	"context"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchClassRoomInfoInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchClassRoomInfoInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchClassRoomInfoInfoLogic {
	return &SearchClassRoomInfoInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchClassRoomInfoInfoLogic) SearchClassRoomInfoInfo(req *types.SearchClassRoomInfoInfoReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
