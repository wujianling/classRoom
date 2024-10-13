package studentDao

import (
	"baize/app/business/student/studentModels"
	"context"
	"github.com/baizeplus/sqly"
)

type IStudentDao interface {
	SelectStudentList(ctx context.Context, db sqly.SqlyContext, student *studentModels.StudentDQL) (studentList []*studentModels.Student, total int64)
	SelectStudentById(ctx context.Context, db sqly.SqlyContext, id uint) (student *studentModels.Student)
	InsertStudent(ctx context.Context, db sqly.SqlyContext, dept *studentModels.Student)
	UpdateStudent(ctx context.Context, db sqly.SqlyContext, dept *studentModels.Student)
	DeleteStudentById(ctx context.Context, db sqly.SqlyContext, id uint)
}
