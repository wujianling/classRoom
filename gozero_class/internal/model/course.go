package model

type Course struct {
	BaseModel
	Name    string `json:"name" gorm:"column:name;not null;comment:名字"`
	Cover   string `json:"cover" gorm:"column:cover;comment:封面"`
	Content string `json:"content" gorm:"column:content;comment:介绍"`
	Status  int    `json:"status" gorm:"column:status;default:1;not null;comment:状态 1-正常 2-停用"`
	Extra   string `json:"extra" gorm:"column:extra;comment:扩展字段"`
}

func (c *Course) TableName() string {
	return "course"
}
