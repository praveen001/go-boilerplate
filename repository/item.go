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

// DeleteByGroupID .
func (r *ItemRepository) DeleteByGroupID(groupID string) error {
	return r.db.Delete(models.Item{}, "playlist_group_id = ?", groupID).Error
}
