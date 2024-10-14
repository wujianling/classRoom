package model

type Teacher struct {
	BaseModel
	Name     string `json:"name" gorm:"column:name;comment:名字"`
	Phone    string `json:"phone" gorm:"column:phone;unique;index;comment:手机号"`
	Password string `json:"password" gorm:"column:password;not null;comment:密码"`
	Avatar   string `json:"avatar" gorm:"column:avatar;comment:头像"`
	Remark   string `json:"remark" gorm:"column:remark;comment:备注"`
	OpenID   string `json:"openID" gorm:"column:open_id;index;comment:微信密钥"`
	CourseID uint   `json:"courseID" gorm:"column:course_id;index;comment:关联课程id"`
	RoomIds  string `json:"roomIds" gorm:"column:room_ids;index;comment:教室ids"`
	Status   int    `json:"status" gorm:"column:status;default:1;not null;comment:状态 1-正常 2-封禁"`
	Extra    string `json:"extra" gorm:"column:extra;comment:扩展字段"`
}

func (t *Teacher) TableName() string {
	return "teacher"
}
