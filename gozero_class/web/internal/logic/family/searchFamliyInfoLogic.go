package family

import (
	"context"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchFamliyInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchFamliyInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchFamliyInfoLogic {
	return &SearchFamliyInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchFamliyInfoLogic) SearchFamliyInfo(req *types.SearchFamilyInfoReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
