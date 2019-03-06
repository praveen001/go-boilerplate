package models

import "time"

// Item ..
type Item struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	// Belongs to Playlist via GroupID
	Playlist        *Playlist `json:"playlist" gorm:"association_foreignkey:GroupID"`
	PlaylistGroupID string    `json:"playlistGroupId"`

	AssetID   string    `json:"assetId"`
	StartTime time.Time `json:"startTime"`
	Duration  uint      `json:"duration"`
}

// TableName .
func (i Item) TableName() string {
	return "new_playlist_items"
}
