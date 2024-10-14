package schoolBanner

import (
	"context"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelSchoolBannerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelSchoolBannerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelSchoolBannerLogic {
	return &DelSchoolBannerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelSchoolBannerLogic) DelSchoolBanner(req *types.DelSchoolBannerReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
