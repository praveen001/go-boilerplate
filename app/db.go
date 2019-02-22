package app

import (
	"github.com/jinzhu/gorm"
	"github.com/praveen001/go-boilerplate/models"
)

// DB .
type DB struct {
	conn *gorm.DB
	User *models.UserService
}

// Use the given connection for db access
func Use(db *gorm.DB) *DB {
	return &DB{
		conn: db,
		User: models.NewUserService(db),
	}
}

// Migrate runs the migrations
func (db *DB) Migrate() {
	db.conn.AutoMigrate()
}

// Close the database connection
func (db *DB) Close() {
	db.conn.Close()
}
