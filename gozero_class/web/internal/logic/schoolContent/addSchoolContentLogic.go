package schoolContent

import (
	"context"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSchoolContentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddSchoolContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSchoolContentLogic {
	return &AddSchoolContentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddSchoolContentLogic) AddSchoolContent(req *types.AddSchoolContentReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
