package model

type Banner struct {
	BaseModel
	Name   string `json:"name" gorm:"column:name;not null;comment:名字"`
	Cover  string `json:"cover" gorm:"column:cover;comment:封面"`
	Link   string `json:"link" gorm:"column:link;comment:链接"`
	Status int    `json:"status" gorm:"column:status;default:1;not null;comment:状态 1-正常 2-停用"`
	Extra  string `json:"extra" gorm:"column:extra;comment:扩展字段"`
}

func (b *Banner) TableName() string {
	return "banner"
}
