package models

import (
	"time"
)

// User ..
// User has many feeds
type User struct {
	ID        uint      `json:"id" gorm:"PRIMARY_KEY"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	Name  string  `json:"name"`
	Email string  `json:"email"`
	Token string  `json:"token" gorm:"COLUMN:authentication_token"`
	Feeds []*Feed `json:"feeds" gorm:"MANY2MANY:feeds_users"`
}
