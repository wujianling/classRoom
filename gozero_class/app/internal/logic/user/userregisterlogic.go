package user

import (
	"context"
	"errors"
	"github.com/smallq_class/internal/model"
	"gorm.io/gorm"

	"github.com/smallq_class/app/internal/svc"
	"github.com/smallq_class/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterReq) (resp *types.BaseResp, err error) {
	if req.Code != "123456" {
		return &types.BaseResp{Code: -1, Msg: "验证码错误，测试验证码：123456"}, nil
	}
	var user model.User
	tx := l.svcCtx.DB.Where(" phone =? ", req.Phone).First(&user)
	if tx.Error != nil && !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return &types.BaseResp{Code: -1, Msg: tx.Error.Error()}, nil
	}
	if user.ID > 0 {
		return &types.BaseResp{Code: -1, Msg: "手机号已注册。"}, nil
	}
	user.Name = req.Name
	user.Phone = req.Phone
	user.OpenID = req.OpenID
	user.Password = req.Password
	user.Avatar = req.Avatar
	result := l.svcCtx.DB.Save(&user)
	if result.Error != nil {
		return &types.BaseResp{Code: -1, Msg: result.Error.Error()}, nil
	}
	return &types.BaseResp{Code: 0}, nil
}
