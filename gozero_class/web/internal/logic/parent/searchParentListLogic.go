package parent

import (
	"context"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchParentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchParentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchParentListLogic {
	return &SearchParentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchParentListLogic) SearchParentList(req *types.SearchParentListReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
