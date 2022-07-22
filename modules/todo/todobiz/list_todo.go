package todobiz

import (
	"API_ToDo_Golang/component"
	"API_ToDo_Golang/modules/todo/todomodel"
	"context"
)

type ListTodoStore interface {
	ListByConditions(ctx context.Context,
		conditions map[string]interface{},
		paging *component.Paging,
		moreKeys ...string) ([]todomodel.ToDoItem, error)
}

type listTodoStore struct {
	store ListTodoStore
}

func NewListTodoStore(store ListTodoStore) *listTodoStore {
	return &listTodoStore{store}
}

func (biz *listTodoStore) ListItems(ctx context.Context,
	conditions map[string]interface{},
	paging *component.Paging,
	moreKeys ...string) ([]todomodel.ToDoItem, error) {

	result, err := biz.store.ListByConditions(ctx, nil, paging)
	if err != nil {
		return nil, err
	}

	return result, err
}
