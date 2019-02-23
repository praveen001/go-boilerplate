package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/praveen001/go-boilerplate/models"
)

// FeedRepository .
type FeedRepository struct {
	db *gorm.DB
}

// NewFeedRepository .
func NewFeedRepository(c *gorm.DB) *FeedRepository {
	return &FeedRepository{c}
}

// New creates a new Feed
func (s *FeedRepository) New(feed *models.Feed) error {
	return s.db.Create(feed).Error
}

// All .
func (s *FeedRepository) All() ([]*models.Feed, error) {
	var f []*models.Feed
	return f, s.db.Preload("User").Find(&f).Error
}
