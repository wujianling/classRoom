// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSysNotice = "sys_notice"

// SysNotice 通知公告表
type SysNotice struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:公告ID" json:"id"` // 公告ID
	Title      string    `gorm:"column:title;not null;comment:公告标题" json:"title"`                // 公告标题
	Type       string    `gorm:"column:type;not null;comment:公告类型（1通知 2公告）" json:"type"`         // 公告类型（1通知 2公告）
	Txt        []byte    `gorm:"column:txt;comment:公告内容" json:"txt"`                             // 公告内容
	DeptID     int64     `gorm:"column:dept_id;comment:发件人所在部门" json:"dept_id"`                  // 发件人所在部门
	DeptIds    string    `gorm:"column:dept_ids;comment:发送部门" json:"dept_ids"`                   // 发送部门
	CreateName string    `gorm:"column:create_name;comment:创建者名称" json:"create_name"`            // 创建者名称
	CreateBy   int64     `gorm:"column:create_by;comment:创建者" json:"create_by"`                  // 创建者
	CreateTime time.Time `gorm:"column:create_time;comment:创建时间" json:"create_time"`             // 创建时间
}

// TableName SysNotice's table name
func (*SysNotice) TableName() string {
	return TableNameSysNotice
}
