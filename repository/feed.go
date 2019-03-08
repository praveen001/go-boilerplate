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
func NewFeedRepository(db *gorm.DB) *FeedRepository {
	return &FeedRepository{db}
}

// FindUserFeed .
func (r *FeedRepository) FindUserFeed(userID, feedID int) (*models.Feed, error) {
	feed := &models.Feed{ID: feedID}

	err := r.db.Raw("SELECT * FROM feeds INNER JOIN feeds_users ON feeds_users.feed_id = feeds.id WHERE feeds_users.user_id = ? AND feeds_users.feed_id = ?", userID, feedID).Scan(feed).Error
	if err != nil {
		return nil, err
	}

	feed.PostFetch()

	return feed, nil
}

// FindUserFeeds .
func (r *FeedRepository) FindUserFeeds(userID int) ([]*models.Feed, error) {
	var feeds []*models.Feed

	rows, err := r.db.Raw("SELECT * FROM feeds INNER JOIN feeds_users ON feeds_users.feed_id = feeds.id WHERE feeds_users.user_id = ?", userID).Rows()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		f := &models.Feed{}
		r.db.ScanRows(rows, f)

		f.PostFetch()

		feeds = append(feeds, f)
	}

	return feeds, nil
}
