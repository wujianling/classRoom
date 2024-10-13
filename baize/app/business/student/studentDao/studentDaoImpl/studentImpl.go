package studentDaoImpl

import (
	"baize/app/business/student/studentModels"
	"context"
	"github.com/baizeplus/sqly"
)

type StudentDao struct {
}

func NewSysUserDao() *StudentDao {
	return &StudentDao{}
}

func (s *StudentDao) SelectStudentList(ctx context.Context, db sqly.SqlyContext, student *studentModels.StudentDQL) (studentList []*studentModels.Student, total int64) {
	whereSql := ``
	if student.Name != "" {
		whereSql += " AND name like concat('%', :name, '%')"
	}
	if student.Sex != 0 {
		whereSql += " AND  sex = :sex"
	}
	if student.Status != "" {
		whereSql += " AND status = :status"
	}

	if whereSql != "" {
		whereSql = " where " + whereSql[4:]
	}
	err := db.NamedSelectPageContext(ctx, &studentList, &total, whereSql, student)
	if err != nil {
		panic(err)
	}
	return
}
func (s *StudentDao) SelectStudentById(ctx context.Context, db sqly.SqlyContext, id uint) (student *studentModels.Student) {
	err := db.GetContext(ctx, &student, "select * from student where id =?", id)
	if err != nil {
		panic(err)
	}
	return
}
func (s *StudentDao) InsertStudent(ctx context.Context, db sqly.SqlyContext, dept *studentModels.Student) {
}
func (s *StudentDao) UpdateStudent(ctx context.Context, db sqly.SqlyContext, dept *studentModels.Student) {
}
func (s *StudentDao) DeleteStudentById(ctx context.Context, db sqly.SqlyContext, id uint) {}
