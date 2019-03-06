package app

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/praveen001/go-boilerplate/models"
	"github.com/praveen001/go-boilerplate/repository"
)

// DB .
type DB struct {
	conn     *gorm.DB
	User     *repository.UserRepository
	Feed     *repository.FeedRepository
	Playlist *repository.PlaylistRepository
}

// initDB initializes DB connections, and prepares all the Models by providing them with the db connection
func (c *Context) initDB() {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@%s/%s?parseTime=true", c.Config.MySQL.User, c.Config.MySQL.Password, c.Config.MySQL.Host, c.Config.MySQL.Database))
	if err != nil {
		c.Logger.Fatal("Unable to connect to database", err.Error())
	}

	c.DB = initRepositories(db)
}

// migrate runs the migrations
func (db *DB) migrate() {
	db.conn.AutoMigrate(models.Playlist{}, models.Item{})
}

// close the database connection
func (db *DB) close() {
	db.conn.Close()
}

// Use the given connection for db access
func initRepositories(db *gorm.DB) *DB {
	return &DB{
		conn:     db,
		User:     repository.NewUserRepository(db),
		Feed:     repository.NewFeedRepository(db),
		Playlist: repository.NewPlaylistRepository(db),
	}
}
