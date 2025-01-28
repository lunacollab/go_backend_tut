package storage

import (
	"Go_backend_tut/modules/item/model"
	"context"
)

func (s *SQLStorage) GetItemById(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error) {
	var data model.TodoItem
	if err := s.db.Where(cond).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
