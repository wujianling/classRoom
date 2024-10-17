package student

import (
	"context"
	"github.com/smallq_class/internal/model"
	"github.com/smallq_class/pkg/utils/role"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddStudentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddStudentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddStudentLogic {
	return &AddStudentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddStudentLogic) AddStudent(req *types.AddStudentReq) (resp *types.BaseResp, err error) {
	quw := l.svcCtx.DB
	adminRole, err := role.GetAdminRole(quw, req.UserID)
	if err != nil {
		return &types.BaseResp{
			Code: -1,
			Msg:  err.Error(),
		}, nil
	}
	isAdmin := role.ContainsRole(adminRole, 1)
	if !isAdmin {
		return &types.BaseResp{
			Code: -1,
			Msg:  "没有权限",
		}, nil
	}
	var student model.Student
	l.svcCtx.DB.Where("phone = ?", req.Phone).First(&student)
	if student.ID > 0 {
		return &types.BaseResp{
			Code: -1,
			Msg:  "手机号已注册",
		}, nil
	}
	student.Name = req.Name
	student.Phone = req.Phone
	student.Sex = req.Sex
	l.svcCtx.DB.Create(&student)

	return &types.BaseResp{
		Code: 0,
		Msg:  "注册成功！",
	}, nil
}
