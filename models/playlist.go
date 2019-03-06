package models

import (
	"time"
)

// PlaylistStatus .
type PlaylistStatus string

// .
const (
	Published PlaylistStatus = "published"
)

// PlaylistType .
type PlaylistType string

// .
const (
	Normal PlaylistType = "normal"
)

// Playlist .
type Playlist struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	// Belongs to Feed
	Feed   *Feed `json:"feed"`
	FeedID uint  `json:"feedId"`

	// Has Items
	Items   []*Item `json:"items"`
	GroupID string  `json:"groupId"`

	PlayOn time.Time      `json:"playOn"`
	Status PlaylistStatus `json:"status"`
	Type   PlaylistType   `json:"type"`
}

// TableName .
func (p Playlist) TableName() string {
	return "new_playlists"
}
