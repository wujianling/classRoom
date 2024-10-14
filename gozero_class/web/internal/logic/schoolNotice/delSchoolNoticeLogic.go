package schoolNotice

import (
	"context"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelSchoolNoticeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelSchoolNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelSchoolNoticeLogic {
	return &DelSchoolNoticeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelSchoolNoticeLogic) DelSchoolNotice(req *types.DelSchoolNoticeReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
