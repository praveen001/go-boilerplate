package app

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// initDB initializes DB connections, and prepares all the Models by providing them with the db connection
func (c *Context) initDB() {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@%s/%s?parseTime=true", c.Config.MySQL.User, c.Config.MySQL.Password, c.Config.MySQL.Host, c.Config.MySQL.Database))
	if c.Config.Environment == Development {
		db.LogMode(true)
	}
	if err != nil {
		c.Logger.Fatal("Unable to connect to database", err.Error())
	}

	c.DB = db
}
