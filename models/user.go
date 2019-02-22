package models

import (
	"github.com/jinzhu/gorm"
)

// User ..
type User struct {
	gorm.Model
	Name string `json:"name"`
}

// UserService .
type UserService struct {
	db *gorm.DB
}

// NewUserService .
func NewUserService(c *gorm.DB) *UserService {
	return &UserService{c}
}

// FindAll .
func (u *UserService) FindAll() ([]*User, error) {
	return nil, nil
}
