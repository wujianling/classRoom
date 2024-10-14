package family

import (
	"context"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchFamilyListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchFamilyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchFamilyListLogic {
	return &SearchFamilyListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchFamilyListLogic) SearchFamilyList(req *types.SearchFamilyListReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
