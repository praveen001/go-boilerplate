package models

import (
	"time"

	"github.com/jinzhu/gorm"
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
	ID        int       `json:"id" gorm:"PRIMARY_KEY"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

	Users     []*User     `json:"-" gorm:"MANY2MANY:feeds_users"`
	Medias    []*Media    `json:"-" gorm:"MANY2MANY:feeds_media"`
	Playlists []*Playlist `json:"-"`
	Account   *Account    `json:"account"`
	AccountID int         `json:"-"`

	Name            string         `json:"name"`
	Code            string         `json:"code" gorm:"COLUMN:channel_code"`
	Timezone        string         `json:"timezone" gorm:"COLUMN:time_zone"`
	InputResolution ResolutionName `json:"-" gorm:"COLUMN:input_video_resolution"`

	FPS int `json:"fps" gorm:"-"`
}

// AfterFind .
func (f *Feed) AfterFind() error {
	// TODO: Find FPS based on Input Resolution
	f.FPS = 25

	f.Timezone = f.Timezone + "(GMT +0530)"

	return nil
}

// Find .
func (f *Feed) Find(db *gorm.DB) error {
	return db.First(f, f).Error
}

// FindUserFeed .
func FindUserFeed(db *gorm.DB, user *User, feed *Feed) error {
	return db.Model(user).Related(feed, "Feeds").Where(feed).Error
}

// FindUserFeeds .
func FindUserFeeds(db *gorm.DB, user *User, feeds *[]*Feed) error {
	return db.Model(user).Related(feeds, "Feeds").Error
}
