package models

import (
	"time"
)

// ResolutionName .
type ResolutionName string

// Resolution .
type Resolution struct {
	Resolution string
	FrameRate  float64
	Interlaced bool
	VideoMode  string
	MaxGFXArea uint
}

//
var (
	HD1080i29_97HZ = Resolution{}
)

// Feed ..
// Feed belongs to many users
// Feed has many playlists
type Feed struct {
	ID        uint      `json:"id" gorm:"PRIMARY_KEY"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

	Name            string         `json:"name"`
	Code            string         `json:"code" gorm:"COLUMN:channel_code"`
	Timezone        string         `json:"timezone" gorm:"COLUMN:time_zone"`
	InputResolution ResolutionName `json:"inputResolution" gorm:"COLUMN:input_video_resolution"`
	Users           []*User        `json:"users" gorm:"MANY2MANY:feeds_users"`
}

// BelongsTo .
func (f *Feed) BelongsTo(userID uint) bool {
	if f == nil {
		return false
	}

	for _, u := range f.Users {
		if u.ID == userID {
			return true
		}
	}

	return false
}
