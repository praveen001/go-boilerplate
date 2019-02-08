package models

import (
	"github.com/jinzhu/gorm"
)

// User ..
type User struct {
	gorm.Model
	Name string `json:"name"`
}

func (db DB) AddUser(u *User) {
	db.db.Get(u.Name)
}
