package models

import (
	"github.com/jinzhu/gorm"
)

// DB is a wrapper over *gorm.DB
// prevents access to gorm.DB methods from controllers
type DB struct {
	db *gorm.DB
}

// Use the given connection for db access
func Use(db *gorm.DB) *DB {
	return &DB{
		db: db,
	}
}

// Migrate runs the migrations
func (db *DB) Migrate() {
	db.db.AutoMigrate(&User{})
}

// Close the database connection
func (db *DB) Close() {
	db.db.Close()
}
