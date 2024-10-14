package model

type CoursePackageLog struct {
	BaseModel
	CoursePID  uint   `json:"coursePID" gorm:"column:course_pid;not null;index;comment:学生课程id"`
	Before     int    `json:"before" gorm:"column:before;not null;comment:变更前"`
	Change     int    `json:"change" gorm:"column:change;not null;comment:变更值"`
	After      int    `json:"after" gorm:"column:after;not null;comment:变更后"`
	Type       int    `json:"type" gorm:"column:type;not null;comment:变更类型 1:增加 2:减少"`
	Memo       string `json:"memo" gorm:"column:memo;comment:备注"`
	OperatorID uint   `json:"operatorID" gorm:"column:operator_id;not null;comment:操作人"`
}

func (c *CoursePackageLog) TableName() string {
	return "course_package_log"
}
