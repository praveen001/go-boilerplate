package app

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/praveen001/go-boilerplate/models"
)

// DB .
type DB struct {
	conn *gorm.DB
	User *models.UserService
}

func (c *Context) initDB() {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@%s/%s?parseTime=true", c.Config.MySQL.User, c.Config.MySQL.Password, c.Config.MySQL.Host, c.Config.MySQL.Database))
	if err != nil {
		c.Logger.Fatal("Unable to connect to database", err.Error())
	}

	c.DB = initServices(db)
}

// migrate runs the migrations
func (db *DB) migrate() {
	db.conn.AutoMigrate()
}

// close the database connection
func (db *DB) close() {
	db.conn.Close()
}

// Use the given connection for db access
func initServices(db *gorm.DB) *DB {
	return &DB{
		conn: db,
		User: models.NewUserService(db),
	}
}
