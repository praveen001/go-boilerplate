package models

import (
	"time"
)

// Feed ..
type Feed struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"-" sql:"index"`
	Name      string     `json:"name"`
	UserID    uint       `json:"userId"`
	User      User       `json:"user"`
}
