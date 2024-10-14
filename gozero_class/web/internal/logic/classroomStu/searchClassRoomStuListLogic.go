package classroomStu

import (
	"context"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchClassRoomStuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchClassRoomStuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchClassRoomStuListLogic {
	return &SearchClassRoomStuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchClassRoomStuListLogic) SearchClassRoomStuList(req *types.SearchClassRoomStuListReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
