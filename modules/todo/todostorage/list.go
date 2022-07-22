package todostorage

import (
	"API_ToDo_Golang/component"
	"API_ToDo_Golang/modules/todo/todomodel"
	"context"
)

func (s *sqlStore) ListByConditions(ctx context.Context,
	conditions map[string]interface{},
	paging *component.Paging,
	moreKeys ...string) ([]todomodel.ToDoItem, error) {
	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	var result []todomodel.ToDoItem

	if err := db.Table(todomodel.ToDoItem{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	db = db.Table(todomodel.ToDoItem{}.TableName()).Where(conditions)
	if err := db.
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
