package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

// DB is a wrapper over *gorm.DB
// prevents access to gorm.DB methods from controllers
type DB struct {
	db *gorm.DB
}

// InitDB connects to the database, and returns *DB
// panics connection failure
func InitDB() *DB {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@%s/%s?parseTime=true", viper.GetString("MYSQL.USER"), viper.GetString("MYSQL.PASSWORD"), viper.GetString("MYSQL.HOST"), viper.GetString("MYSQL.DATABASE")))
	db.DB().SetMaxOpenConns(3)
	if err != nil {
		log.Fatalln("Unable to connect to database", err.Error())
	}

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
