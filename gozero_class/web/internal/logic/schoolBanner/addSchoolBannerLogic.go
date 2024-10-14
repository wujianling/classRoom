package schoolBanner

import (
	"context"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSchoolBannerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddSchoolBannerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSchoolBannerLogic {
	return &AddSchoolBannerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddSchoolBannerLogic) AddSchoolBanner(req *types.AddSchoolBannerReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
