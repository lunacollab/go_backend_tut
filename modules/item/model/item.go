package model

import (
	"Go_backend_tut/common"
	"errors"
)

const (
	EntityName = "Item"
)

var (
	ErrTitleIsBlank = errors.New("title is blank")
	ErrItemDeleted  = errors.New("item is deleted")
)

type TodoItem struct {
	common.SQLModel
	Title       string      `json:"title" gorm:"column:title"`
	Description string      `json:"description" gorm:"column:description"`
	Status      *ItemStatus `json:"status" gorm:"column:status"`
}

type TodoItemCreation struct {
	Id          int         `json:"-" gorm:"column:id"`
	Title       string      `json:"title" gorm:"column:title"`
	Description string      `json:"description" gorm:"column:description"`
	Status      *ItemStatus `json:"status" gorm:"column:status"`
}

type TodoItemUpdate struct {
	Title       *string `json:"title" gorm:"column:title"`
	Description *string `json:"description" gorm:"column:description"`
	Status      *string `json:"status" gorm:"column:status"`
}

func (TodoItemCreation) TableName() string {
	return "todo_items"
}
func (TodoItem) TableName() string {
	return "todo_items"
}
func (TodoItemUpdate) TableName() string {
	return "todo_items"
}
