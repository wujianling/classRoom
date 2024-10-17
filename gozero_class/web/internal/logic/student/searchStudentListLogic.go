package student

import (
	"context"
	"fmt"
	"github.com/smallq_class/internal/model"
	"github.com/smallq_class/pkg/utils/db"
	"github.com/smallq_class/pkg/utils/role"

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
	quw := l.svcCtx.DB
	adminRole, err := role.GetAdminRole(quw, req.UserID)
	fmt.Println("adminRole")
	fmt.Println(adminRole)
	if err != nil {
		return &types.BaseResp{
			Code: -1,
			Msg:  err.Error(),
		}, nil
	}
	isAdmin := role.ContainsRole(adminRole, 1)
	fmt.Println("isAdmin")
	fmt.Println(isAdmin)
	if !isAdmin {
		return &types.BaseResp{
			Code: -1,
			Msg:  "没有权限",
		}, nil
	}
	var students []model.Student
	query := l.svcCtx.DB

	if req.Name != "" {
		query.Where("student.name like ? ", "%"+req.Name+"%")
	}
	if req.Phone != "" {
		query.Where("student.name  like ? ", req.Phone+"%")
	}
	if req.Sex != 0 {
		query.Where("student.name  = ? ", req.Sex)
	}
	if req.Status != 0 {
		query.Where("student.name  = ? ", req.Status)
	}
	query.Joins("left join course_package on ")
	page := db.Paginate(query, req.Page.Page, req.Page.Size)
	if err := query.Find(&students).Error; err != nil {
		return nil, err
	}
	return &types.BaseResp{
		Code: 0,
		Data: struct {
			List []model.Student `json:"list"`
			Page db.ReturnPage   `json:"page"`
		}{students, page},
	}, nil
}
