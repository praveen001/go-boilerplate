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
	ID        uint      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	Name            string         `json:"name"`
	Code            string         `json:"code" gorm:"column:channel_code"`
	Timezone        string         `json:"timezone" gorm:"column:time_zone"`
	InputResolution ResolutionName `json:"inputResolution" gorm:"column:input_video_resolution"`
	Users           []*User        `json:"users" gorm:"many2many:feeds_users"`
	Playlists       []*Playlist    `json:"playlists"`
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
