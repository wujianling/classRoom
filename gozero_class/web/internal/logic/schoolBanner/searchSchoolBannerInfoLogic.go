package schoolBanner

import (
	"context"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchSchoolBannerInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchSchoolBannerInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchSchoolBannerInfoLogic {
	return &SearchSchoolBannerInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchSchoolBannerInfoLogic) SearchSchoolBannerInfo(req *types.SearchSchoolBannerInfoReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
