package role

import (
	"errors"
	"github.com/smallq_class/dao/model"
	"gorm.io/gorm"
	"sort"
)

func GetAdminRole(db *gorm.DB, userId int64) (role []int64, err error) {
	var userRole []model.SysUserRole
	db.Select("role_id").Where("user_id = ?", userId).Find(&userRole)
	if len(userRole) == 0 {
		return nil, errors.New("管理员未找到")
	}
	var roleList []int64
	for _, v := range userRole {

		roleList = append(roleList, v.RoleID)
	}
	return roleList, nil
}
func ContainsRole(s []int64, e int) bool {
	result := make([]int, len(s))
	for i, v := range s {
		result[i] = int(v)
	}
	// 对切片进行排序
	sort.Ints(result)

	// 使用二分查找来查找元素
	index := sort.SearchInts(result, e)

	// 如果找到了元素，那么index不会等于len(s)
	return index < len(result) && result[index] == e
}
