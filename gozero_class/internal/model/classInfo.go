package model

import "time"

type ClassInfo struct {
	BaseModel
	CourseID    uint      `json:"courseID" gorm:"column:course_id;not null;index;comment:关联课程id"`
	ClassRoomID uint      `json:"classRoomID" gorm:"column:classroom_id;not null;index;comment:教室id"`
	TeacherID   uint      `json:"teacherID" gorm:"column:teacher_id;not null;index;comment:关联老师id"`
	MaxNum      int       `json:"maxNum" gorm:"column:max_num;not null;comment:最大人数"`
	CurNum      int       `json:"curNum" gorm:"column:cur_num;not null;default:0;comment:当前人数"`
	Status      int       `json:"status" gorm:"column:status;default:0;not null;comment:状态 0-待发布 1-未开始 2已取消 3已完成 4 销课"`
	StartTime   time.Time `json:"startTime" gorm:"column:start_time;not null;comment:开始时间"`
	EndTime     time.Time `json:"endTime" gorm:"column:end_time;not null;comment:结束时间"`
	CheckTime   time.Time `json:"checkTime" gorm:"column:check_time;comment:签到时间"`
	Duration    int       `json:"duration" gorm:"column:duration;not null;comment:时长分钟"`
	Extra       string    `json:"extra" gorm:"column:extra;comment:扩展字段"`
}

func (c *ClassInfo) TableName() string {
	return "class_info"
}
