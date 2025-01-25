package bussiness

import (
	"Go_backend_tut/modules/item/model"
	"context"
	"strings"
)

type CreateItemStorage interface {
	CreateItem(ctx context.Context, data *model.TodoItemCreation) error
}
type createItemBussiness struct {
	storage CreateItemStorage
}

func NewCreateItemBussiness(storage CreateItemStorage) *createItemBussiness {
	return &createItemBussiness{storage: storage}
}

func (bussiness *createItemBussiness) CreateNewItem(ctx context.Context, data *model.TodoItemCreation) error {
	title := strings.TrimSpace(data.Title)
	if title == "" {
		return model.ErrTitleIsBlank
	}

	if err := bussiness.storage.CreateItem(ctx, data); err != nil {
		return err
	}
	return nil
}
