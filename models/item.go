package models

import "time"

// ItemStatus .
type ItemStatus string

//
const (
	ItemStatusPresent                ItemStatus = "Present"
	ItemStatusLive                              = "Live"
	ItemStatusInvalid                           = "Invalid"
	ItemStatusProcessing                        = "Processing"
	ItemStatusNotPresent                        = "Not Present"
	ItemStatusSegmentAbsent                     = "Segment Absent"
	ItemStatusIncorrectSOM                      = "Incorrect SOM"
	ItemStatusDurationMismatch                  = "Duration Mismatch"
	ItemStatusSmallDuration                     = "Small Duration"
	ItemStatusInvalidArguments                  = "Invalid Arguments"
	ItemStatusFileOffsetAbsent                  = "File Offset Absent"
	ItemStatusIncorrectStartTime                = "Incorrect Start Time"
	ItemStatusSubtitleAbsent                    = "Subtitle Absent"
	ItemStatusHole                              = "Hole"
	ItemStatusSkipped                           = "Skipped"
	ItemStatusRescue                            = "Rescue"
	ItemStatusDisable                           = "Disable"
	ItemStatusInvalidType                       = "Invalid Type"
	ItemStatusInvalidDynamicAsset               = "Invalid Dynamic Asset"
	ItemStatusInvalidDynamicTemplate            = "Invalid Dynamic Template"
	ItemStatusInvalidDynamcParams               = "Invalid Dynamic Params"
	ItemStatusInvalidSCTEParams                 = "Invalid SCTE Params"
	ItemStatusInvalidLiveParams                 = "Invalid Live Trigger Params"
	ItemStatusDynamicAssetAbsent                = "Dynamic Asset Absent"
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

	AssetID     string      `json:"assetId"`
	Title       string      `json:"title"`
	SegmentID   int         `json:"segmentId"`
	StartTime   int64       `json:"startTime"`
	Duration    int         `json:"duration"`
	Locked      bool        `json:"locked"`
	Status      ItemStatus  `json:"status"`
	ItemType    ItemType    `json:"itemType"`
	AssetStatus MediaStatus `json:"assetStatus"`
}

// TableName .
func (i Item) TableName() string {
	return "new_playlist_items"
}
