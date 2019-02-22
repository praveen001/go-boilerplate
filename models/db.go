package models

import (
	"github.com/jinzhu/gorm"
)

// DB is a wrapper over *gorm.DB
// prevents access to gorm.DB methods from controllers
type DB struct {
	db *gorm.DB
}
