package teacher

import (
	"context"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchTeacherListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchTeacherListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchTeacherListLogic {
	return &SearchTeacherListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchTeacherListLogic) SearchTeacherList(req *types.SearchTeacherListReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
