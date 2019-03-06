package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/praveen001/go-boilerplate/models"
)

// PlaylistRepository .
type PlaylistRepository struct {
	db *gorm.DB
}

// NewPlaylistRepository .
func NewPlaylistRepository(c *gorm.DB) *PlaylistRepository {
	return &PlaylistRepository{c}
}

// Save .
func (r *PlaylistRepository) Save(playlist *models.Playlist) error {
	return r.db.Save(playlist).Error
}

// Find .
func (r *PlaylistRepository) Find(playlistID uint) (*models.Playlist, error) {
	p := &models.Playlist{}
	return p, r.db.Preload("Items").First(p, playlistID).Error
}

// FindByDate .
func (r *PlaylistRepository) FindByDate(date uint64) ([]*models.Playlist, error) {
	var p []*models.Playlist
	return p, r.db.Find(&p, models.Playlist{
		PlayOn: date,
	}).Error
}
