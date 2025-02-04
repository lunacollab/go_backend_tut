package storage

import (
	"Go_backend_tut/common"
	"Go_backend_tut/modules/item/model"
	"context"
)

func (s *SQLStorage) CreateItem(ctx context.Context, data *model.TodoItemCreation) error {
	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
