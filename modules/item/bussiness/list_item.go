package bussiness

import (
	"Go_backend_tut/common"
	"Go_backend_tut/modules/item/model"
	"context"
)

type ListItemStorage interface {
	ListItem(ctx context.Context, filter *model.Filter, paging *common.Paging, moreKeys ...string) ([]model.TodoItem, error)
}
type ListItemBussiness struct {
	storage ListItemStorage
}

func NewListItemBussiness(storage ListItemStorage) *ListItemBussiness {
	return &ListItemBussiness{storage: storage}
}

func (bussiness *ListItemBussiness) ListItemBuz(ctx context.Context, filter *model.Filter, paging *common.Paging) ([]model.TodoItem, error) {

	data, err := bussiness.storage.ListItem(ctx, filter, paging)
	if err != nil {
		if err == common.RecordNotFound {
			return common.ErrCannotGetEntity(model.EntityName, err)
		}
		return nil, common.ErrCannotListEntity(model.EntityName, err)
	}
	return data, nil
}
