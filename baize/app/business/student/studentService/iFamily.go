package studentService

import (
	"baize/app/business/student/studentModels"
	"github.com/gin-gonic/gin"
)

type IFamilyService interface {
	SelectFamilyList(c *gin.Context, family *studentModels.Family) (list []*studentModels.Family)
	AddFamilyMember(c *gin.Context, family *studentModels.Family)
	UpdateFamilyMember(c *gin.Context, family *studentModels.Family)
	DelFamilyMember(c *gin.Context, id uint)
}
