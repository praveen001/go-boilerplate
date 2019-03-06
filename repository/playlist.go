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

// New .
func (r *PlaylistRepository) New(playlist *models.Playlist) error {
	r.db.Create(playlist)
	return r.db.Save(playlist).Error
}

// Find .
func (r *PlaylistRepository) Find(playlistID uint) (*models.Playlist, error) {
	p := &models.Playlist{}
	return p, r.db.Preload("Items").First(p, playlistID).Error
}
