package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/praveen001/go-boilerplate/models"
)

// UserRepository .
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository .
func NewUserRepository(c *gorm.DB) *UserRepository {
	return &UserRepository{c}
}

// New creates a new user
func (s *UserRepository) New(user *models.User) error {
	return s.db.Create(user).Error
}

// All .
func (s *UserRepository) All() ([]*models.User, error) {
	var u []*models.User
	return u, s.db.Preload("Feeds").Find(&u).Error
}

// DeleteAll .
func (s *UserRepository) DeleteAll() error {
	return s.db.Unscoped().Delete(models.User{}).Error
}

// Find .
func (s *UserRepository) Find(ID uint) (*models.User, error) {
	var u models.User
	return &u, s.db.First(&u, ID).Error
}

// Delete .
func (s *UserRepository) Delete(ID uint) error {
	return s.db.Delete(&models.User{}, ID).Error
}

// Update ..
func (s *UserRepository) Update(u *models.User) error {
	return s.db.Model(&u).Update(u).Error
}

func (s *UserRepository) preload() *gorm.DB {
	// Add following struct tag on fields that shouldn't be preloaded
	// `gorm:"PRELOAD:false"`
	return s.db.Set("gorm:auto_preload", true)
}
