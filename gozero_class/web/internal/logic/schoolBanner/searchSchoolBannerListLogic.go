package schoolBanner

import (
	"context"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchSchoolBannerListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchSchoolBannerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchSchoolBannerListLogic {
	return &SearchSchoolBannerListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchSchoolBannerListLogic) SearchSchoolBannerList(req *types.SearchSchoolBannerListReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
