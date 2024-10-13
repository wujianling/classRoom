package studentModels

import "baize/app/baize"

type Parent struct {
	baize.BaseEntity
	Name     string `json:"name" db:"name"`
	Phone    string `json:"phone" db:"phone"`
	Password string `json:"password" db:"password"`
	Avatar   string `json:"avatar" db:"avatar"`
	Remark   string `json:"remark" db:"remark"`
	OpenID   string `json:"openID" db:"open_id"`
	Status   string `json:"status" db:"status"`
	Extra    string `json:"extra" db:"extra"`
}
