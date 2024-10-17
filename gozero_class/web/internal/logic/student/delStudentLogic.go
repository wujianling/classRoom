package student

import (
	"context"
	"github.com/smallq_class/internal/model"
	"github.com/smallq_class/pkg/utils/role"

	"github.com/smallq_class/web/internal/svc"
	"github.com/smallq_class/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelStudentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelStudentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelStudentLogic {
	return &DelStudentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelStudentLogic) DelStudent(req *types.DelStudentReq) (resp *types.BaseResp, err error) {
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
	l.svcCtx.DB.Delete(&student)
	return &types.BaseResp{
		Code: 0,
		Msg:  "删除成功！",
	}, nil
}
