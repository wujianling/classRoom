package studentService

import (
	"baize/app/business/student/studentModels"
	"github.com/gin-gonic/gin"
)

type IStudentService interface {
	SelectStudentList(c *gin.Context, student *studentModels.Student) (list []*studentModels.Student)
	InsertStudent(c *gin.Context, student *studentModels.Student)
	UpdateStudent(c *gin.Context, student *studentModels.Student)
	DelStudent(c *gin.Context, student *studentModels.Student)
}
