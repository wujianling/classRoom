package teacher

import (
	"context"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchTeacherInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchTeacherInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchTeacherInfoLogic {
	return &SearchTeacherInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchTeacherInfoLogic) SearchTeacherInfo(req *types.SearchTeacherInfoReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
