package models

import "time"

// ItemGroup .
type ItemGroup struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	// Belongs to playlist
	Playlist   *Playlist `json:"playlist"`
	PlaylistID uint      `json:"playlistId"`

	// Has many items
	Items []*Item `json:"items"`
}
