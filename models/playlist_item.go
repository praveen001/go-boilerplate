package models

import "time"

// Item ..
type Item struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	// Belongs to ItemGroup
	ItemGroup   *ItemGroup `json:"itemGroup"`
	ItemGroupID uint       `json:"itemGroupId"`

	AssetID   string    `json:"assetId"`
	StartTime time.Time `json:"startTime"`
	Duration  uint      `json:"duration"`
}
