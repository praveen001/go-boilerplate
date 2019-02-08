package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

type DB struct {
	db *gorm.DB
}

func InitDB() *DB {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@%s/%s?parseTime=true", viper.GetString("MYSQL.USER"), viper.GetString("MYSQL.PASSWORD"), viper.GetString("MYSQL.HOST"), viper.GetString("MYSQL.DATABASE")))
	if err != nil {
		log.Fatalln("Unable to connect to database", err.Error())
	}

	db.AutoMigrate(&User{})

	return &DB{
		db: db,
	}
}

func (db *DB) Close() {
	db.db.Close()
}
