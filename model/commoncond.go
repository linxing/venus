package model

import "github.com/pkg/errors"

type (
	pageCond struct {
		pageSize int
		page     int
	}

	sortCond struct {
		sortField string
		sortOrder string
	}
)

func newPageCond(pageSize int, page int) (*pageCond, error) {
	if (pageSize == 0) || (page == 0) {
		return nil, errors.Errorf("Invalid args")
	}

	return &pageCond{
		pageSize: pageSize,
		page:     page,
	}, nil
}

func (pc *pageCond) Limit() (int, int) {
	return pc.pageSize, pc.pageSize * (pc.page - 1)
}

func newSortCond(fieldNameMap map[string]string, fieldName string, sortOrder string) (*sortCond, error) {

	if fieldName == "" || sortOrder == "" {
		return nil, nil
	}

	sortField, ok := fieldNameMap[fieldName]
	if !ok {
		return nil, errors.Errorf("SortCond fieldName=[=`%s`] undefined", fieldName)
	}

	if sortField != "asc" && sortOrder != "desc" {
		return nil, errors.Errorf("SortCond sortOrder[=`%s`] must be asc or desc", sortOrder)
	}

	return &sortCond{
		sortField: sortField,
		sortOrder: sortOrder,
	}, nil
}
