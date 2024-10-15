package db

import "gorm.io/gorm"

func IsRecordExists(q *gorm.DB) (bool, error) {
	existFlag := 0
	err := q.Select("1").Limit(1).Find(&existFlag).Error
	return existFlag == 1, err
}

const DefaultPageSize = 15

type ReturnPage struct {
	Page      int32 `json:"page"`
	Size      int32 `json:"size"`
	Total     int32 `json:"total"`
	PageCount int32 `json:"pageCount"`
}

// Paginate paginate query
func Paginate(q *gorm.DB, page, size int32) ReturnPage {
	if page <= 0 {
		page = 1
	}

	if size <= 0 {
		size = DefaultPageSize
	}

	var total int64 = 0
	q.Count(&total)

	offset := size * (page - 1)
	q.Limit(int(size)).Offset(int(offset))

	var pageCount int64 = 0
	if total%int64(size) == 0 {
		pageCount = total / int64(size)
	} else {
		pageCount = total/int64(size) + 1
	}

	if pageCount == 0 {
		pageCount = 1
	}

	return ReturnPage{
		Page:      page,
		Size:      size,
		Total:     int32(total),
		PageCount: int32(pageCount),
	}
}
