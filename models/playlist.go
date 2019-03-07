package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// PlaylistStatus .
type PlaylistStatus string

// .
const (
	PlaylistStatusPublished PlaylistStatus = "published"
	PlaylistStatusNew                      = "new"
)

// PlaylistType .
type PlaylistType string

// .
const (
	PlaylistTypeNormal PlaylistType = "normal"
)

// Playlist .
type Playlist struct {
	ID        int       `json:"id" gorm:"PRIMARY_KEY"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

	// Belongs to Feed
	Feed   *Feed `json:"feed"`
	FeedID int   `json:"-"`

	// Has Items
	Items []*Item `json:"items"`

	PlayOn int            `json:"playOn"`
	Status PlaylistStatus `json:"status"`
	Type   PlaylistType   `json:"type"`
}

// TableName .
func (p Playlist) TableName() string {
	return "new_playlists"
}

// Create .
func (p *Playlist) Create(db *gorm.DB) error {
	return db.Create(p).Error
}

// Find .
func (p *Playlist) Find(db *gorm.DB) error {
	return db.Preload("Items").Find(p).Error
}

// Delete .
func (p *Playlist) Delete(db *gorm.DB) error {
	return db.Delete(p).Error
}

// FindPlaylistByDate .
func FindPlaylistByDate(db *gorm.DB, date int, feedID int) ([]*Playlist, error) {
	var playlists []*Playlist

	return playlists, db.Find(&playlists, Playlist{
		PlayOn: date,
		FeedID: feedID,
	}).Error
}
