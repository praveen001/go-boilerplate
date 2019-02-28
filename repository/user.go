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

// Find .
func (s *UserRepository) Find(ID uint) (*models.User, error) {
	var u models.User
	return &u, s.db.First(&u, ID).Error
}

// FindByToken .
func (s *UserRepository) FindByToken(tok string) (*models.User, error) {
	var u models.User
	return &u, s.db.Preload("Feeds").Where("token = ?", tok).First(&u).Error
}

func (s *UserRepository) preload() *gorm.DB {
	// Add following struct tag on fields that shouldn't be preloaded
	// `gorm:"PRELOAD:false"`
	return s.db.Set("gorm:auto_preload", true)
}
