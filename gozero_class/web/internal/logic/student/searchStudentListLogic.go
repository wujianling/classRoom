package student

import (
	"context"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchStudentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchStudentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchStudentListLogic {
	return &SearchStudentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchStudentListLogic) SearchStudentList(req *types.SearchStudentListReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
