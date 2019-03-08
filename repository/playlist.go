package repository

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/praveen001/go-boilerplate/models"
)

// PlaylistRepository .
type PlaylistRepository struct {
	db *gorm.DB
}

// NewPlaylistRepository .
func NewPlaylistRepository(db *gorm.DB) *PlaylistRepository {
	return &PlaylistRepository{db}
}

// Create .
func (r *PlaylistRepository) Create(playlist *models.Playlist) error {
	return r.db.Create(playlist).Error
}

// Read .
func (r *PlaylistRepository) Read(playlistID int) (*models.Playlist, error) {
	p := &models.Playlist{
		ID: playlistID,
	}

	return p, r.db.Preload("Items").Find(p).Error
}

// Delete .
func (r *PlaylistRepository) Delete(playlist *models.Playlist) error {
	return r.db.Delete(playlist).Error
}

// FindPlaylistByDate .
func (r *PlaylistRepository) FindPlaylistByDate(feedID int, date time.Time) ([]*models.Playlist, error) {
	var playlists []*models.Playlist

	return playlists, r.db.Find(&playlists, models.Playlist{
		PlayOn: date,
		FeedID: feedID,
	}).Error
}
