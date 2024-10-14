package model

type ClassInfoStu struct {
	BaseModel
	ClassInfoID uint   `json:"classInfoID" gorm:"column:class_info_id;not null;index;comment:课程id"`
	StudentID   uint   `json:"studentID" gorm:"column:student_id;not null;index;comment:关联学生id"`
	Status      int    `json:"status" gorm:"column:status;default:1;not null;comment:状态 1-待签到 2已签到 3已取消"`
	Cover       string `json:"cover" gorm:"column:cover;comment:图片"`
	Video       string `json:"video" gorm:"column:video;comment:视频"`
	Extra       string `json:"extra" gorm:"column:extra;comment:扩展字段"`
}

func (c *ClassInfoStu) TableName() string {
	return "class_info_stu"
}
