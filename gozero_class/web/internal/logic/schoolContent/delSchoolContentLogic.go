package schoolContent

import (
	"context"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelSchoolContentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelSchoolContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelSchoolContentLogic {
	return &DelSchoolContentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelSchoolContentLogic) DelSchoolContent(req *types.DelSchoolContentReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
