package schoolContent

import (
	"context"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveSchoolContentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveSchoolContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveSchoolContentLogic {
	return &SaveSchoolContentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveSchoolContentLogic) SaveSchoolContent(req *types.SaveSchoolContentReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
