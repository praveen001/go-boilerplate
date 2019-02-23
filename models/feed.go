package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Feed ..
type Feed struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"-" sql:"index"`
	Name      string     `json:"name"`
	UserID    uint       `json:"userId"`
	User      User       `json:"user"`
}

// FeedService .
type FeedService struct {
	db *gorm.DB
}

// NewFeedService .
func NewFeedService(c *gorm.DB) *FeedService {
	return &FeedService{c}
}

// New creates a new Feed
func (s *FeedService) New(feed *Feed) error {
	return s.db.Create(feed).Error
}

// All .
func (s *FeedService) All() ([]*Feed, error) {
	var f []*Feed
	return f, s.db.Preload("User").Find(&f).Error
}
