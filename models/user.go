package models

import (
	"time"
)

// User ..
// User has many feeds
type User struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	Name  string  `json:"name"`
	Email string  `json:"email"`
	Token string  `json:"token" gorm:"column:authentication_token"`
	Feeds []*Feed `json:"feeds" gorm:"many2many:feeds_users"`
}
