package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/praveen001/go-boilerplate/models"
)

// ItemRepository .
type ItemRepository struct {
	db *gorm.DB
}

// NewItemRepository .
func NewItemRepository(c *gorm.DB) *ItemRepository {
	return &ItemRepository{c}
}

// DeleteMulti .
func (r *ItemRepository) DeleteMulti(items []*models.Item) error {
	ids := make([]uint, len(items))
	for i, item := range items {
		ids[i] = item.ID
	}
	return r.db.Where("id in (?)", ids).Delete(models.Item{}).Error
}
