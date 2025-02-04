package bussiness

import (
	"Go_backend_tut/common"
	"Go_backend_tut/modules/item/model"
	"context"
	"strings"
)

type CreateItemStorage interface {
	CreateItem(ctx context.Context, data *model.TodoItemCreation) error
}
type CreateItemBussiness struct {
	storage CreateItemStorage
}

func NewCreateItemBussiness(storage CreateItemStorage) *CreateItemBussiness {
	return &CreateItemBussiness{storage: storage}
}

func (bussiness *CreateItemBussiness) CreateNewItem(ctx context.Context, data *model.TodoItemCreation) error {
	title := strings.TrimSpace(data.Title)
	if title == "" {
		return model.ErrTitleIsBlank
	}

	if err := bussiness.storage.CreateItem(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(model.EntityName, err)
	}
	return nil
}
