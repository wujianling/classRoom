package classroomStu

import (
	"context"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchClassRoomStuInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchClassRoomStuInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchClassRoomStuInfoLogic {
	return &SearchClassRoomStuInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchClassRoomStuInfoLogic) SearchClassRoomStuInfo(req *types.SearchClassRoomStuInfoReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
