package models

import (
	"time"

	"github.com/jinzhu/gorm"
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

// UserService .
type UserService struct {
	db *gorm.DB
}

// NewUserService .
func NewUserService(c *gorm.DB) *UserService {
	return &UserService{c}
}

// New creates a new user
func (s *UserService) New(user *User) error {
	return s.db.Create(user).Error
}

// All .
func (s *UserService) All() ([]*User, error) {
	var u []*User
	return u, s.db.Preload("Feeds").Find(&u).Error
}

// DeleteAll .
func (s *UserService) DeleteAll() error {
	return s.db.Unscoped().Delete(User{}).Error
}

// Find .
func (s *UserService) Find(ID uint) (*User, error) {
	var u User
	return &u, s.db.First(&u, ID).Error
}

// Delete .
func (s *UserService) Delete(ID uint) error {
	return s.db.Delete(&User{}, ID).Error
}

// Update ..
func (s *UserService) Update(u *User) error {
	return s.db.Model(&u).Update(u).Error
}

func (s *UserService) preload() *gorm.DB {
	// Add following struct tag on fields that shouldn't be preloaded
	// `gorm:"PRELOAD:false"`
	return s.db.Set("gorm:auto_preload", true)
}
