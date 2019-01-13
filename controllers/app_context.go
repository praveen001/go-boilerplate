package controllers

import (
	"github.com/jinzhu/gorm"
)

// AppContext holds the context for each request
// Everything in context must be thread-safe
type AppContext struct {
	DB *gorm.DB
}
