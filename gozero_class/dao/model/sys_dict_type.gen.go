// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSysDictType = "sys_dict_type"

// SysDictType 字典类型表
type SysDictType struct {
	DictID     int64     `gorm:"column:dict_id;primaryKey;autoIncrement:true;comment:字典主键" json:"dict_id"` // 字典主键
	DictName   string    `gorm:"column:dict_name;comment:字典名称" json:"dict_name"`                           // 字典名称
	DictType   string    `gorm:"column:dict_type;comment:字典类型" json:"dict_type"`                           // 字典类型
	Status     string    `gorm:"column:status;default:0;comment:状态（0正常 1停用）" json:"status"`                // 状态（0正常 1停用）
	CreateBy   int64     `gorm:"column:create_by;comment:创建者" json:"create_by"`                            // 创建者
	CreateTime time.Time `gorm:"column:create_time;comment:创建时间" json:"create_time"`                       // 创建时间
	UpdateBy   int64     `gorm:"column:update_by;comment:更新者" json:"update_by"`                            // 更新者
	UpdateTime time.Time `gorm:"column:update_time;comment:更新时间" json:"update_time"`                       // 更新时间
	Remark     string    `gorm:"column:remark;comment:备注" json:"remark"`                                   // 备注
}

// TableName SysDictType's table name
func (*SysDictType) TableName() string {
	return TableNameSysDictType
}
