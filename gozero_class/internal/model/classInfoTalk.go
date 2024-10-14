package model

type ClassInfoTalk struct {
	BaseModel
	ClassInfoID uint   `json:"classInfoID" gorm:"column:class_info_id;not null;index;comment:课程id"`
	StudentID   uint   `json:"studentID" gorm:"column:student_id;not null;index;comment:关联学生id"`
	TeacherID   uint   `json:"teacherID" gorm:"column:teacher_id;not null;index;comment:关联老师id"`
	Content     string `json:"content" gorm:"column:content;not null;comment:内容"`
	Type        uint   `json:"type" gorm:"column:type;not null;comment:类型"`
	FromTea     bool   `json:"fromTea" gorm:"column:from_tea;not null;comment:是否老师发送"`
	IsStuRead   bool   `json:"isStuRead" gorm:"column:is_stu_read;not null;comment:家长是否阅读"`
	IsTeaRead   bool   `json:"isTeaRead" gorm:"column:is_tea_read;not null;comment:老师是否阅读"`
	Extra       string `json:"extra" gorm:"column:extra;comment:扩展字段"`
}

func (c *ClassInfoTalk) TableName() string {
	return "class_info_talk"
}
