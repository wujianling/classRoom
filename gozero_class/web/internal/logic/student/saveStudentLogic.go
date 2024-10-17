package student

import (
	"context"
	"github.com/smallq_class/internal/model"
	"github.com/smallq_class/pkg/utils/role"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveStudentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveStudentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveStudentLogic {
	return &SaveStudentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveStudentLogic) SaveStudent(req *types.SaveStudentReq) (resp *types.BaseResp, err error) {
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
	l.svcCtx.DB.Where("id = ?", req.StudentID).First(&student)
	if student.ID == 0 {
		return &types.BaseResp{
			Code: -1,
			Msg:  "学生未找到",
		}, nil
	}

	if req.Name != "" {
		student.Name = req.Name
	}
	if req.Sex != 0 {
		student.Sex = req.Sex
	}
	if req.Phone != "" {
		var student2 model.Student
		l.svcCtx.DB.Where("phone = ?", req.Phone).First(&student2)
		if student.ID > 0 {
			return &types.BaseResp{
				Code: -1,
				Msg:  "手机号已注册",
			}, nil
		}
		student.Phone = req.Phone
	}
	l.svcCtx.DB.Save(&student)
	return &types.BaseResp{
		Code: 0,
		Msg:  "修改成功！",
	}, nil
}
