package studentModels

import "baize/app/baize"

type Student struct {
	baize.BaseEntity
	Name     string `json:"name" db:"name"`
	Sex      int    `json:"sex" db:"sex"`
	Birthday string `json:"birthday" db:"birthday"`
	Status   string `json:"status" db:"status"`
	Extra    string `json:"extra" db:"extra"`
}

type StudentDQL struct {
	baize.BaseEntityDQL
	Name   string `json:"name" db:"name"`
	Sex    int    `json:"sex" db:"sex"`
	Status string `json:"status" db:"status"`
}
