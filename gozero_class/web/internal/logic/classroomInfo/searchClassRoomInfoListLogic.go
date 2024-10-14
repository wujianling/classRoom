package classroomInfo

import (
	"context"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchClassRoomInfoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchClassRoomInfoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchClassRoomInfoListLogic {
	return &SearchClassRoomInfoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchClassRoomInfoListLogic) SearchClassRoomInfoList(req *types.SearchClassRoomInfoListReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
