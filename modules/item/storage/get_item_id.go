package storage

import (
	"Go_backend_tut/common"
	"Go_backend_tut/modules/item/model"
	"context"
	"gorm.io/gorm"
)

func (s *SQLStorage) GetItemById(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error) {
	var data model.TodoItem
	if err := s.db.Where(cond).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &data, nil
}
