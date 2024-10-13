package studentService

import (
	"baize/app/business/student/studentModels"
	"github.com/gin-gonic/gin"
)

type IParentService interface {
	SelectParentList(c *gin.Context, parent *studentModels.Parent) (list []*studentModels.Parent)
	InsertParent(c *gin.Context, parent *studentModels.Parent)
	UpdateParent(c *gin.Context, parent *studentModels.Parent)
	DelParent(c *gin.Context, parent *studentModels.Parent)
}
