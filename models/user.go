package models

import (
	"time"
)

// User ..
type User struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"-" sql:"index"`
	Name      string     `json:"name"`
	Feeds     []Feed     `json:"feeds"`
}
