package models

import "time"

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
// Playlist belongs to Feed
type Playlist struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	FeedID    uint           `json:"feedId"`
	Feed      Feed           `json:"feed"`
	StartTime time.Time      `json:"startTime"`
	Status    PlaylistStatus `json:"status"`
	Type      PlaylistType   `json:"type"`
}
