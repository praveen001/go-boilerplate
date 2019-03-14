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
func NewItemRepository(db *gorm.DB) *ItemRepository {
	return &ItemRepository{db}
}

// BatchDelete .
func (r *ItemRepository) BatchDelete(ids []int) error {
	return r.db.Delete(models.Item{}, "id IN (?)", ids).Error
	// return r.db.Raw("DELETE FROM new_playlist_items WHERE id IN (?)", ids).Error
}
