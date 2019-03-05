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

// Find .
func (s *FeedRepository) Find(feedID uint) (*models.Feed, error) {
	f := &models.Feed{}
	return f, s.db.Preload("Users").First(f, feedID).Error
}
