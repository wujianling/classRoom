package model

type ClassRoom struct {
	BaseModel
	CourseID  uint   `json:"courseID" gorm:"column:course_id;not null;index;comment:关联课程id"`
	Name      string `json:"name" gorm:"column:name;not null;comment:名字"`
	Introduce string `json:"introduce" gorm:"column:introduce;comment:介绍"`
	MaxNum    int    `json:"maxNum" gorm:"column:max_num;not null;comment:最大人数"`
	Duration  int    `json:"duration" gorm:"column:duration;not null;comment:时长分钟"`
	Cover     string `json:"cover" gorm:"column:cover;comment:封面"`
	StuList   string `json:"stuList" gorm:"column:stu_ist;comment:学生数组字符串"`
	Status    int    `json:"status" gorm:"column:status;default:1;not null;comment:状态 1-正常 2-停用"`
	Extra     string `json:"extra" gorm:"column:extra;comment:扩展字段"`
}

func (c *ClassRoom) TableName() string {
	return "class_room"
}
