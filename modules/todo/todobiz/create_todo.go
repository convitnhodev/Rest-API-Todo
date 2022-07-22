package todobiz

import (
	 "API_ToDo_Golang/modules/todo/todomodel"
	 "context"
	 "errors"
	 "strings"
)

type CreateTodoStore interface {
	 Create(ctx context.Context, data *todomodel.ToDoItemCreate) error
}

type createTodoStore struct {
	 store CreateTodoStore
}

func NewCreateTodoStore(store CreateTodoStore) *createTodoStore {
	 return &createTodoStore{store}
}

func (biz *createTodoStore) CreateTodo(ctx context.Context, data *todomodel.ToDoItemCreate) error {
	 data.Title = strings.TrimSpace(data.Title)
	 
	 if data.Title == "" {
		  return errors.New("title cannot be blank")
	 }
	 
	 return biz.store.Create(ctx, data)
}
