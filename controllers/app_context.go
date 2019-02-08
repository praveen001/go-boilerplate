package controllers

import (
	"github.com/praveen001/go-boilerplate/models"
)

// AppContext holds the context for each request
// Everything in context must be thread-safe
type AppContext struct {
	DB *models.DB
}
