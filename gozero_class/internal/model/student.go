package model

type Student struct {
	BaseModel
	Phone  string `json:"phone" gorm:"column:phone;size:20;not null;uniqueIndex;comment:注册手机号"`
	Name   string `json:"name" gorm:"column:name;not null;index;comment:名字"`
	Sex    int    `json:"sex" gorm:"column:sex;default:3;comment:性别 1男 2女 3未填写"`
	Status int    `json:"status" gorm:"column:status;default:1;not null;comment:状态 1-正常 2-结束"`
	Extra  string `json:"extra" gorm:"column:extra;comment:扩展字段"`
}

func (s *Student) TableName() string {
	return "student"
}
