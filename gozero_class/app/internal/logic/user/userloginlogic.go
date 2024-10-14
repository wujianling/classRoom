package user

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/smallq_class/internal/model"
	"time"

	"github.com/smallq_class/app/internal/svc"
	"github.com/smallq_class/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.UserLoginReq) (resp *types.BaseResp, err error) {
	var user model.User
	tx := l.svcCtx.DB.Where("open_id = ? ", req.OpenID).First(&user)
	if tx.Error != nil {
		return &types.BaseResp{Code: -1, Msg: tx.Error.Error()}, nil
	}
	if user.ID == 0 {
		return &types.BaseResp{Code: -1, Msg: "微信未关联账号。"}, nil
	}
	authConfig := l.svcCtx.Config.Auth
	now := time.Now().Unix()
	accessToken, err := l.GenToken(now, authConfig.AccessSecret, map[string]interface{}{"uid": user.ID, "username": user.Name}, authConfig.AccessExpire)
	if err != nil {
		return &types.BaseResp{Code: -1, Msg: err.Error()}, nil
	}
	return &types.BaseResp{Code: 0, Data: accessToken}, nil
}
func (l *UserLoginLogic) GenToken(iat int64, secretKey string, payloads map[string]interface{}, seconds int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	for k, v := range payloads {
		claims[k] = v
	}
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
