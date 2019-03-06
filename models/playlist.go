package models

import (
	"time"
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
	ID        uint      `json:"id" gorm:"PRIMARY_KEY"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

	// Belongs to Feed
	Feed   *Feed `json:"feed"`
	FeedID uint  `json:"-"`

	// Has Items
	Items []*Item `json:"items"`

	PlayOn uint64         `json:"playOn"`
	Status PlaylistStatus `json:"status"`
	Type   PlaylistType   `json:"type"`
}

// TableName .
func (p Playlist) TableName() string {
	return "new_playlists"
}
