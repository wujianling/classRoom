package user

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/smallq_class/app/internal/svc"
	"github.com/smallq_class/app/internal/types"
	"github.com/smallq_class/internal/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.BaseResp, err error) {
	uid, err := l.ctx.Value("uid").(json.Number).Int64()
	fmt.Println(uid)
	var user model.User
	tx := l.svcCtx.DB.Where("id = ? ", uid).First(&user)
	if tx.Error != nil {
		return &types.BaseResp{Code: -1, Msg: tx.Error.Error()}, nil
	}
	if user.ID == 0 {
		return &types.BaseResp{Code: -1, Msg: "信息获取失败。"}, nil
	}
	user.Password = ""
	user.OpenID = ""
	return &types.BaseResp{Code: 0, Data: user}, nil
}
