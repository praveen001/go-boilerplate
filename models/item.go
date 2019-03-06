package models

import "time"

// ItemStatus .
type ItemStatus string

//
const (
	ItemStatusOk ItemStatus = "ok"
)

// ItemType .
type ItemType string

//
const (
	ItemTypePrimary              ItemType = "primary"
	ItemTypeSecondary                     = "secondary"
	ItemTypeIndependentSecondary          = "independentSecondary"
	ItemTypeHole                          = "hole"
)

// Item ..
type Item struct {
	ID        uint      `json:"id" gorm:"PRIMARY_KEY"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

	// Belongs to Playlist via GroupID
	Playlist        *Playlist `json:"-" gorm:"ASSOCIATION_FOREIGNKEY:GroupID"`
	PlaylistGroupID string    `json:"-"`

	AssetID   string     `json:"assetId"`
	Title     string     `json:"title"`
	SegmentID uint       `json:"segmentId"`
	StartTime uint       `json:"startTime"`
	Duration  uint       `json:"duration"`
	Locked    bool       `json:"locked"`
	Status    ItemStatus `json:"status"`
	ItemType  ItemType   `json:"itemType"`
}

// TableName .
func (i Item) TableName() string {
	return "new_playlist_items"
}
