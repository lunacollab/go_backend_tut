package bussiness

import (
	"Go_backend_tut/modules/item/model"
	"context"
)

type DeleteItemStorageById interface {
	GetItemById(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error)
	DeleteItemById(ctx context.Context, cond map[string]interface{}) error
}
type DeleteItemBussinessById struct {
	storage DeleteItemStorageById
}

func NewDeleteItemBussinessById(storage DeleteItemStorageById) *DeleteItemBussinessById {
	return &DeleteItemBussinessById{storage: storage}
}

func (bussiness *DeleteItemBussinessById) DeleteItemBuzById(ctx context.Context, id int) error {

	data, err := bussiness.storage.GetItemById(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	if data.Status != nil && *data.Status == model.ItemStatusDeleted {
		return model.ErrItemDeleted
	}

	if err := bussiness.storage.DeleteItemById(ctx, map[string]interface{}{"id": id}); err != nil {
		return err
	}

	return nil
}
