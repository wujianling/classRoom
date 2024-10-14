package model

type User struct {
	BaseModel
	Name     string `json:"name" gorm:"column:name;index;comment:名字"`
	Role     int    `json:"role" gorm:"column:role;index;default:3;comment:角色 1 管理员 2销售 3其他"`
	Phone    string `json:"phone" gorm:"column:phone;unique;index;comment:手机号"`
	Password string `json:"password" gorm:"column:password;not null;comment:密码"`
	Avatar   string `json:"avatar" gorm:"column:avatar;comment:头像"`
	Remark   string `json:"remark" gorm:"column:remark;comment:备注"`
	OpenID   string `json:"openID" gorm:"column:open_id;index;comment:微信密钥"`
	Status   int    `json:"status" gorm:"column:status;default:1;not null;comment:状态 1-正常 2-结束"`
	Extra    string `json:"extra" gorm:"column:extra;comment:扩展字段"`
}

func (u *User) TableName() string {
	return "user"
}
