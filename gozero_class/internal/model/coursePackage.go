package model

type CoursePackage struct {
	BaseModel
	CourseID     uint   `json:"courseID" gorm:"column:course_id;not null;index;comment:关联课程id"`
	ClassRoomID  uint   `json:"classRoomID" gorm:"column:class_room_id;not null;index;comment:关联教室id"`
	StudentID    uint   `json:"studentID" gorm:"column:student_id;not null;index;comment:关联学生id"`
	OrderID      string `json:"orderID" gorm:"column:order_id;not null;index;comment:合同id"`
	FirstOrderID string `json:"firstOrderID" gorm:"column:first_order_id;not null;default:-1;index;comment:初始合同id"`
	Issue        int    `json:"issue" gorm:"column:issue;not null;default:1;index;comment:第几期"`
	FromType     int    `json:"fromType" gorm:"column:from_type;not null;default:1;index;comment:1 自然流 2 老师介绍 3学生介绍"`
	FromRef      uint   `json:"fromRef" gorm:"column:from_ref;not null;default:0;index;comment:关联人员id"`
	Status       int    `json:"status" gorm:"column:status;not null;default:1;comment:状态 1正常 2停用"`
	Num          int    `json:"num" gorm:"column:num;not null;default:0;comment:剩余课时"`
}

func (c *CoursePackage) TableName() string {
	return "course_package"
}
