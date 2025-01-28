package storage

import (
	"Go_backend_tut/modules/item/model"
	"context"
)

func (s *SQLStorage) DeleteItemById(ctx context.Context, cond map[string]interface{}) error {
	if err := s.db.Table(model.TodoItem{}.TableName()).Where(cond).Updates(map[string]interface{}{
		"status": "Deleted",
	}).Error; err != nil {
		return err
	}
	return nil
}
