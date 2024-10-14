package schoolBanner

import (
	"context"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveSchoolBannerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveSchoolBannerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveSchoolBannerLogic {
	return &SaveSchoolBannerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveSchoolBannerLogic) SaveSchoolBanner(req *types.SaveSchoolBannerReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
