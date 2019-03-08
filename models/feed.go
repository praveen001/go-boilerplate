package models

import (
	"time"
)

// Feed ..
// Feed belongs to many users
// Feed has many playlists
type Feed struct {
	ID        int       `json:"id" gorm:"PRIMARY_KEY"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

	Users     []*User     `json:"-" gorm:"MANY2MANY:feeds_users"`
	Medias    []*Media    `json:"-" gorm:"MANY2MANY:feeds_media"`
	Playlists []*Playlist `json:"-"`
	Account   *Account    `json:"account,omitempty"`
	AccountID int         `json:"-"`

	Name            string `json:"name"`
	Code            string `json:"code" gorm:"COLUMN:channel_code"`
	Timezone        string `json:"timezone" gorm:"COLUMN:time_zone"`
	InputResolution string `json:"-" gorm:"COLUMN:input_video_resolution"`

	FPS float64 `json:"fps" gorm:"-"`
}

// PostFetch .
func (f *Feed) PostFetch() {
	f.FPS = ResolutionMap[f.InputResolution].FrameRate

	// TODO: Need to fix this
	f.Timezone = f.Timezone + "(GMT +0530)"
}
