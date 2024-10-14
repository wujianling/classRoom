package schoolContent

import (
	"context"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchSchoolContentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchSchoolContentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchSchoolContentListLogic {
	return &SearchSchoolContentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchSchoolContentListLogic) SearchSchoolContentList(req *types.SearchSchoolContentListReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
