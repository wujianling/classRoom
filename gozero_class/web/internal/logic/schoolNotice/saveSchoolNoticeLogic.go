package schoolNotice

import (
	"context"
	"github.com/smallq_class/pkg/utils/role"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveSchoolNoticeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveSchoolNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveSchoolNoticeLogic {
	return &SaveSchoolNoticeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveSchoolNoticeLogic) SaveSchoolNotice(req *types.SaveSchoolNoticeReq) (resp *types.BaseResp, err error) {
	quw := l.svcCtx.DB
	adminRole, err := role.GetAdminRole(quw, req.UserID)
	if err != nil {
		return &types.BaseResp{
			Code: -1,
			Msg:  err.Error(),
		}, nil
	}
	isAdmin := role.ContainsRole(adminRole, 1)
	if !isAdmin {
		return &types.BaseResp{
			Code: -1,
			Msg:  "没有权限",
		}, nil
	}

	return
}
