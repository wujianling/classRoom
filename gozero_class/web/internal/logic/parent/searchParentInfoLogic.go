package parent

import (
	"context"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchParentInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchParentInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchParentInfoLogic {
	return &SearchParentInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchParentInfoLogic) SearchParentInfo(req *types.SearchParentInfoReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
