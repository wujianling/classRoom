package model

type Notice struct {
	BaseModel
	Content string `json:"content" gorm:"column:content;comment:内容"`
	Status  int    `json:"status" gorm:"column:status;default:1;not null;comment:状态 1-正常 2-停用"`
	Extra   string `json:"extra" gorm:"column:extra;comment:扩展字段"`
}

func (n *Notice) TableName() string {
	return "notice"
}
