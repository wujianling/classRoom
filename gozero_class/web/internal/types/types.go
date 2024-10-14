// Code generated by goctl. DO NOT EDIT.
package types

type AddClassRoomInfoReq struct {
	UserID      int64 `json:"userId"`
	CourseID    uint  `json:"courseID"`
	ClassRoomID uint  `json:"classRoomId"`
	TeacherID   uint  `json:"teacherID"`
	MaxNum      int   `json:"maxNum"`
	CurNum      int   `json:"curNum"`
	Status      int   `json:"status"`
	StartTime   int64 `json:"startTime"`
	EndTime     int64 `json:"endTime"`
	Duration    int   `json:"duration"`
}

type AddClassRoomReq struct {
	UserID    int64  `json:"userId"`
	CourseID  uint   `json:"courseID"`
	Name      string `json:"name"`
	Introduce string `json:"introduce"`
	MaxNum    int    `json:"maxNum"`
	StuList   string `json:"stuList"`
	Duration  int    `json:"duration"`
	Cover     string `json:"cover"`
	Status    int    `json:"status"`
}

type AddClassRoomStuReq struct {
	UserID      int64 `json:"userId"`
	ClassInfoID uint  `json:"classInfoID"`
	StudentID   uint  `json:"studentID"`
}

type AddClassRoomTalkReq struct {
	UserID      int64  `json:"userId"`
	ClassInfoID uint   `json:"classInfoID"`
	StudentID   uint   `json:"studentID"`
	TeacherID   uint   `json:"teacherID"`
	Content     string `json:"content"`
	Type        uint   `json:"type"`
	FromTea     bool   `json:"fromTea"`
	IsStuRead   bool   `json:"isStuRead"`
	IsTeaRead   bool   `json:"isTeaRead"`
}

type AddCourseLogReq struct {
	UserID     int64  `json:"userId"`
	CoursePID  uint   `json:"coursePID"`
	Before     int    `json:"before"`
	Change     int    `json:"change"`
	After      int    `json:"after"`
	Type       int    `json:"type"`
	Memo       string `json:"memo,optional"`
	OperatorID uint   `json:"operatorID"`
}

type AddCoursePackageReq struct {
	UserID       int64  `json:"userId"`
	CourseID     uint   `json:"courseId"`
	ClassRoomID  uint   `json:"classRoomID"`
	StudentID    uint   `json:"studentID"`
	OrderID      string `json:"orderID"`
	Issue        int    `json:"issue"`
	FromType     int    `json:"fromType"`
	FromRef      uint   `json:"fromRef"`
	Status       int    `json:"status"`
	FirstOrderID string `json:"firstOrderID"`
	Num          int    `json:"num"`
}

type AddCourseReq struct {
	UserID  int64  `json:"userId"`
	Name    string `json:"name"`
	Cover   string `json:"cover"`
	Content string `json:"content"`
}

type AddFamilyReq struct {
	UserID    int64  `json:"userId"`
	StudentID uint   `json:"studentId"`
	ParentID  uint   `json:"parentId"`
	Remark    string `json:"remark,optional"`
}

type AddParentReq struct {
	UserID   int64  `json:"userId"`
	Phone    string `json:"phone"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Remark   string `json:"remark,optional"`
}

type AddSchoolBannerReq struct {
}

type AddSchoolContentReq struct {
}

type AddSchoolNoticeReq struct {
}

type AddStudentReq struct {
	UserID int64  `json:"userId"`
	Phone  string `json:"phone"`
	Name   string `json:"name"`
	Sex    int    `json:"sex"`
}

type AddTeacherReq struct {
	UserID   int64  `json:"userId"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Avatar   string `json:"avatar,optional"`
	Remark   string `json:"remark,optional"`
	Typer    int    `json:"typer"`
	CourseID uint   `json:"courseID"`
	RoomIds  string `json:"roomIds,optional"`
}

type BaseResp struct {
	Code string      `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type DelClassRoomInfoReq struct {
	UserID          int64 `json:"userId"`
	ClassRoomInfoID uint  `json:"classRoomInfoId"`
}

type DelClassRoomReq struct {
	UserID      int64 `json:"userId"`
	ClassRoomID uint  `json:"classRoomId"`
}

type DelClassRoomStuReq struct {
	UserID         int64 `json:"userId"`
	ClassRoomStuID uint  `json:"classRoomStuId"`
}

type DelClassRoomTalkReq struct {
	UserID          int64 `json:"userId"`
	ClassInfoTalkID uint  `json:"classInfoTalkID"`
}

type DelCourseLogReq struct {
	UserID      int64 `json:"userId"`
	CourseLogID uint  `json:"courseLogId"`
}

type DelCoursePackageReq struct {
	UserID          int64 `json:"userId"`
	CoursePackageID uint  `json:"coursePackageID"`
}

type DelCourseReq struct {
	UserID   int64 `json:"userId"`
	CourseID uint  `json:"courseId"`
}

type DelFamilyReq struct {
	UserID   int64 `json:"userId"`
	FamilyID uint  `json:"familyId"`
}

type DelParentReq struct {
	UserID   int64 `json:"userId"`
	ParentID uint  `json:"parentId"`
}

type DelSchoolBannerReq struct {
}

type DelSchoolContentReq struct {
}

type DelSchoolNoticeReq struct {
}

type DelStudentReq struct {
	UserID    int64 `json:"userId"`
	StudentID uint  `json:"studentId"`
}

type DelTeacherReq struct {
	UserID    int64 `json:"userId"`
	TeacherID int64 `json:"teacherID"`
}

type ReturnPage struct {
	Page      int32 `json:"page"`
	Size      int32 `json:"size"`
	Total     int32 `json:"total"`
	PageCount int32 `json:"pageCount"`
}

type SaveClassRoomInfoReq struct {
	UserID          int64 `json:"userId"`
	ClassRoomInfoID uint  `json:"classRoomInfoId"`
	CourseID        uint  `json:"courseID,optional"`
	ClassRoomID     uint  `json:"classRoomId,optional"`
	TeacherID       uint  `json:"teacherID,optional"`
	MaxNum          int   `json:"maxNum,optional"`
	CurNum          int   `json:"curNum,optional"`
	Status          int   `json:"status,optional"`
	StartTime       int64 `json:"startTime,optional"`
	EndTime         int64 `json:"endTime,optional"`
	Duration        int   `json:"duration,optional"`
}

type SaveClassRoomReq struct {
	UserID      int64  `json:"userId"`
	ClassRoomID uint   `json:"classRoomId"`
	CourseID    uint   `json:"courseID,optional"`
	Name        string `json:"name,optional"`
	Introduce   string `json:"introduce,optional"`
	MaxNum      int    `json:"maxNum,optional"`
	StuList     string `json:"stuList,optional"`
	Duration    int    `json:"duration,optional"`
	Cover       string `json:"cover,optional"`
	Status      int    `json:"status,optional"`
}

type SaveClassRoomStuReq struct {
	UserID         int64  `json:"userId"`
	ClassRoomStuID uint   `json:"classRoomStuId"`
	ClassInfoID    uint   `json:"classInfoID,optional"`
	StudentID      uint   `json:"studentID,optional"`
	Status         int    `json:"status,optional"`
	Cover          string `json:"cover,optional"`
	Video          string `json:"video,optional"`
}

type SaveClassRoomTalkReq struct {
	UserID          int64 `json:"userId"`
	ClassInfoTalkID uint  `json:"classInfoTalkID"`
	ClassInfoID     uint  `json:"classInfoID,optional"`
	StudentID       uint  `json:"studentID,optional"`
	TeacherID       uint  `json:"teacherID,optional"`
	Type            uint  `json:"type,optional"`
	FromTea         bool  `json:"fromTea,optional"`
	IsStuRead       bool  `json:"isStuRead,optional"`
	IsTeaRead       bool  `json:"isTeaRead,optional"`
}

type SaveCourseLogReq struct {
	UserID      int64  `json:"userId"`
	CourseLogID uint   `json:"courseLogId"`
	CoursePID   uint   `json:"coursePID,optional"`
	Before      int    `json:"before,optional"`
	Change      int    `json:"change,optional"`
	After       int    `json:"after,optional"`
	Type        int    `json:"type,optional"`
	Memo        string `json:"memo,optional"`
	OperatorID  uint   `json:"operatorID,optional"`
}

type SaveCoursePackageReq struct {
	UserID          int64  `json:"userId"`
	CoursePackageID uint   `json:"coursePackageID"`
	CourseID        uint   `json:"courseId,optional"`
	ClassRoomID     uint   `json:"classRoomID,optional"`
	StudentID       uint   `json:"studentID,optional"`
	OrderID         string `json:"orderID,optional"`
	Issue           int    `json:"issue,optional"`
	FromType        int    `json:"fromType,optional"`
	FromRef         uint   `json:"fromRef,optional"`
	Status          int    `json:"status,optional"`
	FirstOrderID    string `json:"firstOrderID,optional"`
	Num             int    `json:"num,optional"`
}

type SaveCourseReq struct {
	UserID   int64  `json:"userId"`
	CourseID uint   `json:"courseId"`
	Name     string `json:"name,optional"`
	Cover    string `json:"cover,optional"`
	Content  string `json:"content,optional"`
}

type SaveFamilyReq struct {
	UserID   int64  `json:"userId"`
	FamilyID uint   `json:"familyId"`
	Remark   string `json:"remark"`
}

type SaveParentReq struct {
	UserID   int64  `json:"userId"`
	ParentID uint   `json:"parentId"`
	Phone    string `json:"phone,optional"`
	Name     string `json:"name,optional"`
	Status   int    `json:"status，optional"`
	Password string `json:"password,optional"`
	Remark   string `json:"remark,optional"`
}

type SaveSchoolBannerReq struct {
}

type SaveSchoolContentReq struct {
}

type SaveSchoolNoticeReq struct {
}

type SaveStudentReq struct {
	UserID    int64  `json:"userId"`
	StudentID uint   `json:"studentId"`
	Phone     string `json:"phone,optional"`
	Name      string `json:"name,optional"`
	Sex       int    `json:"sex，optional"`
	Status    int    `json:"status，optional"`
}

type SaveTeacherReq struct {
	UserID    int64  `json:"userId"`
	TeacherID int64  `json:"teacherID"`
	Name      string `json:"name,optional"`
	Phone     string `json:"phone,optional"`
	Password  string `json:"password,optional"`
	Avatar    string `json:"avatar,optional"`
	Remark    string `json:"remark,optional"`
	Typer     int    `json:"typer,optional"`
	CourseID  uint   `json:"courseID,optional"`
	RoomIds   string `json:"roomIds,optional"`
}

type SearchClassRoomInfoInfoReq struct {
	UserID          int64 `json:"userId"`
	ClassRoomInfoID uint  `json:"classRoomInfoId"`
}

type SearchClassRoomInfoListReq struct {
	UserID      int64      `json:"userId"`
	ClassRoomID uint       `json:"classRoomId,optional"`
	CourseID    uint       `json:"courseID,optional"`
	TeacherID   uint       `json:"teacherID,optional"`
	MaxNum      int        `json:"maxNum,optional"`
	CurNum      int        `json:"curNum,optional"`
	Status      int        `json:"status,optional"`
	StartTime1  int64      `json:"startTime1,optional"`
	StartTime2  int64      `json:"startTime2,optional"`
	EndTime1    int64      `json:"endTime1,optional"`
	EndTime2    int64      `json:"endTime2,optional"`
	Duration    int        `json:"duration"`
	Page        SelectPage `json:"page,optional"`
}

type SearchClassRoomInfoReq struct {
	UserID      int64 `json:"userId"`
	ClassRoomID uint  `json:"classRoomId"`
}

type SearchClassRoomListReq struct {
	UserID   int64      `json:"userId"`
	Page     SelectPage `json:"page,optional"`
	CourseID uint       `json:"courseID,optional"`
	Name     string     `json:"name,optional"`
	Stu      string     `json:"stu"`
	Status   int        `json:"status,optional"`
}

type SearchClassRoomStuInfoReq struct {
	UserID         int64 `json:"userId"`
	ClassRoomStuID uint  `json:"classRoomStuId"`
}

type SearchClassRoomStuListReq struct {
	UserID      int64      `json:"userId"`
	ClassRoomID uint       `json:"classRoomId,optional"`
	StudentID   uint       `json:"studentID,optional"`
	Status      int        `json:"status,optional"`
	Page        SelectPage `json:"page,optional"`
}

type SearchClassRoomTalkInfoReq struct {
	UserID          int64 `json:"userId"`
	ClassInfoTalkID uint  `json:"classInfoTalkID"`
}

type SearchClassRoomTalkListReq struct {
	UserID      int64      `json:"userId"`
	ClassInfoID uint       `json:"classInfoID,optional"`
	StudentID   uint       `json:"studentID,optional"`
	TeacherID   uint       `json:"teacherID,optional"`
	Type        uint       `json:"type,optional"`
	FromTea     bool       `json:"fromTea,optional"`
	IsStuRead   bool       `json:"isStuRead,optional"`
	IsTeaRead   bool       `json:"isTeaRead,optional"`
	Page        SelectPage `json:"page,optional"`
}

type SearchCourseInfoReq struct {
	UserID   int64 `json:"userId"`
	CourseID uint  `json:"courseId"`
}

type SearchCourseListReq struct {
	UserID int64 `json:"userId"`
}

type SearchCourseLogInfoReq struct {
	UserID      int64 `json:"userId"`
	CourseLogID uint  `json:"courseLogId"`
}

type SearchCourseLogListReq struct {
	UserID     int64      `json:"userId"`
	CoursePID  uint       `json:"coursePID"`
	Type       int        `json:"type"`
	OperatorID uint       `json:"operatorID"`
	Page       SelectPage `json:"page,optional"`
}

type SearchCoursePackageInfoReq struct {
	UserID          int64 `json:"userId"`
	CoursePackageID uint  `json:"coursePackageID"`
}

type SearchCoursePackageListReq struct {
	UserID      int64      `json:"userId"`
	CourseID    uint       `json:"courseId,optional"`
	ClassRoomID uint       `json:"classRoomID,optional"`
	StudentID   uint       `json:"studentID,optional"`
	OrderID     string     `json:"orderID，optional"`
	Issue       int        `json:"issue,optional"`
	FromType    int        `json:"fromType,optional"`
	FromRef     uint       `json:"fromRef,optional"`
	Status      int        `json:"status,optional"`
	Page        SelectPage `json:"page,optional"`
}

type SearchFamilyInfoReq struct {
	UserID    int64 `json:"userId"`
	StudentID uint  `json:"studentId"`
	ParentID  uint  `json:"parentId"`
}

type SearchFamilyListReq struct {
	UserID    int64      `json:"userId"`
	StudentID uint       `json:"studentId,optional"`
	ParentID  uint       `json:"parentId,optional"`
	Page      SelectPage `json:"page,optional"`
}

type SearchParentInfoReq struct {
	UserID   int64 `json:"userId"`
	ParentID uint  `json:"parentId"`
}

type SearchParentListReq struct {
	UserID    int64      `json:"userId"`
	Phone     string     `json:"phone,optional"`
	StudentID uint       `json:"studentId,optional"`
	Name      string     `json:"name,optional"`
	Status    int        `json:"status，optional"`
	Page      SelectPage `json:"page,optional"`
}

type SearchSchoolBannerInfoReq struct {
}

type SearchSchoolBannerListReq struct {
}

type SearchSchoolContentInfoReq struct {
}

type SearchSchoolContentListReq struct {
}

type SearchSchoolNoticeInfoReq struct {
}

type SearchSchoolNoticeListReq struct {
}

type SearchStudentInfoReq struct {
	UserID    int64 `json:"userId"`
	StudentID uint  `json:"studentId"`
}

type SearchStudentListReq struct {
	UserID int64      `json:"userId"`
	Phone  string     `json:"phone,optional"`
	Name   string     `json:"name,optional"`
	Sex    int        `json:"sex，optional"`
	Status int        `json:"status，optional"`
	Page   SelectPage `json:"page,optional"`
}

type SearchTeacherInfoReq struct {
	UserID    int64 `json:"userId"`
	TeacherID int64 `json:"teacherID"`
}

type SearchTeacherListReq struct {
	UserID   int64      `json:"userId"`
	Name     string     `json:"name,optional"`
	Phone    string     `json:"phone,optional"`
	Typer    int        `json:"typer,optional"`
	CourseID uint       `json:"courseID,optional"`
	RoomIds  string     `json:"roomIds,optional"`
	Page     SelectPage `json:"page,optional"`
}

type SelectPage struct {
	Page int32 `json:"page,optional"`
	Size int32 `json:"size,optional"`
}
