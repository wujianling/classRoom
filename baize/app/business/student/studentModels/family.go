package studentModels

import "baize/app/baize"

type Family struct {
	baize.BaseEntity
	StudentID uint   `json:"studentID" db:"student_id"`
	ParentID  uint   `json:"parentID" db:"parent_id"`
	Remark    string `json:"remark" db:"remark"`
	Status    string `json:"status" db:"status"`
	Extra     string `json:"extra" db:"extra"`
}
