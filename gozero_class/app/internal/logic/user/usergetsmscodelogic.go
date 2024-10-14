package user

import (
	"context"
	"github.com/smallq_class/pkg/utils/sms"

	"github.com/smallq_class/app/internal/svc"
	"github.com/smallq_class/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserGetSMSCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserGetSMSCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserGetSMSCodeLogic {
	return &UserGetSMSCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserGetSMSCodeLogic) UserGetSMSCode(req *types.UserGetSMSCodeReq) (resp *types.BaseResp, err error) {
	code := "123456"
	err = sms.SendSms(req.Phone, req.Mtehod, code)
	if err != nil {
		return &types.BaseResp{Code: -1}, err
	}
	return &types.BaseResp{Code: 0, Data: code}, nil
}
