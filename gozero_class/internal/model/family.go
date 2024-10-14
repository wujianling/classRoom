package model

type Family struct {
	BaseModel
	StudentID uint   `json:"studentID" gorm:"column:student_id;not null;comment:学生id"`
	ParentID  uint   `json:"parentID" gorm:"column:parent_id;not null;comment:家长id"`
	Remark    string `json:"remark" gorm:"column:remark;comment:备注"`
	Extra     string `json:"extra" gorm:"column:extra;comment:扩展字段"`
}

func (f *Family) TableName() string {
	return "family"
}
