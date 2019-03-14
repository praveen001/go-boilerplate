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
	ID        int       `json:"id" gorm:"PRIMARY_KEY"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

	// Belongs to Playlist
	Playlist   *Playlist `json:"-"`
	PlaylistID int       `json:"-"`

	AssetID   string     `json:"assetId"`
	Title     string     `json:"title"`
	SegmentID int        `json:"segmentId"`
	StartTime int64        `json:"startTime"`
	Duration  int        `json:"duration"`
	Locked    bool       `json:"locked"`
	Status    ItemStatus `json:"status"`
	ItemType  ItemType   `json:"itemType"`
}

// TableName .
func (i Item) TableName() string {
	return "new_playlist_items"
}
