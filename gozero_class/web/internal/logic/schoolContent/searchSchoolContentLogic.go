package schoolContent

import (
	"context"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchSchoolContentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchSchoolContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchSchoolContentLogic {
	return &SearchSchoolContentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchSchoolContentLogic) SearchSchoolContent(req *types.SearchSchoolContentInfoReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
