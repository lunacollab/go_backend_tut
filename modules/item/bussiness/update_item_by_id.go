package bussiness

import (
	"Go_backend_tut/common"
	"Go_backend_tut/modules/item/model"
	"context"
)

type UpdateItemStorageById interface {
	GetItemById(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error)
	UpdateItemById(ctx context.Context, cond map[string]interface{}, dataUpdate *model.TodoItemUpdate) error
}
type UpdateItemBussinessById struct {
	storage UpdateItemStorageById
}

func NewUpdateItemBussinessById(storage UpdateItemStorageById) *UpdateItemBussinessById {
	return &UpdateItemBussinessById{storage: storage}
}

func (bussiness *UpdateItemBussinessById) UpdateItemBuzById(ctx context.Context, id int, dataUpdate *model.TodoItemUpdate) error {

	data, err := bussiness.storage.GetItemById(ctx, map[string]interface{}{"id": id})
	if err != nil {
		if err == common.RecordNotFound {
			return common.ErrCannotGetEntity(model.EntityName, err)
		}
		return common.ErrCannotUpdateEntity(model.EntityName, err)
	}
	if data.Status != nil && *data.Status == model.ItemStatusDeleted {
		return common.ErrEntityDeleted(model.EntityName, model.ErrItemDeleted)
	}

	if err := bussiness.storage.UpdateItemById(ctx, map[string]interface{}{"id": id}, dataUpdate); err != nil {
		return common.ErrCannotUpdateEntity(model.EntityName, err)
	}

	return nil
}
