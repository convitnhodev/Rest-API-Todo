package todostorage

import (
	 "API_ToDo_Golang/modules/todo/todomodel"
	 "context"
)

func (s *sqlStore) Create(ctx context.Context, data *todomodel.ToDoItemCreate) error {
	 db := s.db
	 if err := db.Create(data).Error; err != nil {
		  return err
	 }
	 return nil
}
