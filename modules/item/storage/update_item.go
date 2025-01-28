package storage

import (
	"Go_backend_tut/modules/item/model"
	"context"
)

func (s *SQLStorage) UpdateItemById(ctx context.Context, cond map[string]interface{}, dataUpdate *model.TodoItemUpdate) error {
	if err := s.db.Where(cond).Updates(dataUpdate).Error; err != nil {
		return err
	}
	return nil
}
