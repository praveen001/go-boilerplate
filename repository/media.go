package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/praveen001/go-boilerplate/models"
)

// MediaRepository .
type MediaRepository struct {
	db *gorm.DB
}

// NewMediaRepository .
func NewMediaRepository(c *gorm.DB) *MediaRepository {
	return &MediaRepository{c}
}

// FindByFeedID .
func (r *MediaRepository) FindByFeedID(feedID uint) ([]*models.Media, error) {
	var m []*models.Media
	return m, r.db.Find(&m, models.Media{
		ID: feedID,
	}).Error
}

// DB .
func (r *MediaRepository) DB() *gorm.DB {
	return r.db
}
