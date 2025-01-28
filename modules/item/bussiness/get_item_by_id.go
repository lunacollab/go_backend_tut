package bussiness

import (
	"Go_backend_tut/modules/item/model"
	"context"
)

type GetItemStorageById interface {
	GetItemById(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error)
}
type GetItemBussinessById struct {
	storage GetItemStorageById
}

func NewGetItemBussinessById(storage GetItemStorageById) *GetItemBussinessById {
	return &GetItemBussinessById{storage: storage}
}

func (bussiness *GetItemBussinessById) GetItemBuzById(ctx context.Context, id int) (*model.TodoItem, error) {

	data, err := bussiness.storage.GetItemById(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}
	return data, nil
}
