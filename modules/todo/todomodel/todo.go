package todomodel

import (
	"time"
)

type ToDoItem struct {
	Id        int        `json:"id" gorm:"column:id"`
	Title     string     `json:"title" gorm:"column:title"`
	Detail    string     `json:"detail" gorm:"column:detail"`
	Status    string     `json:"status" gorm:"column:status;default:Doing"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

func (t ToDoItem) TableName() string {
	return "todo_items"
}
