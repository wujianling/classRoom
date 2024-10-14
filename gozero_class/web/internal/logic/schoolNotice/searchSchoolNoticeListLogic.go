package schoolNotice

import (
	"context"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchSchoolNoticeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchSchoolNoticeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchSchoolNoticeListLogic {
	return &SearchSchoolNoticeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchSchoolNoticeListLogic) SearchSchoolNoticeList(req *types.SearchSchoolNoticeListReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
