package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"

	"github.com/praveen001/go-boilerplate/models"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/praveen001/go-boilerplate/controllers"

	"github.com/praveen001/go-boilerplate/router"

	"github.com/spf13/viper"
)

func init() {
	configPath := flag.String("configPath", "./", "Config file path")
	flag.Parse()

	viper.AddConfigPath(*configPath)
	viper.SetConfigName("config")
	if viper.ReadInConfig() != nil {
		log.Fatalln("Error reading config file")
	}
}

func main() {
	// Database
	db := initDB()
	defer db.Close()

	// Redis
	redisPool := initRedis()
	defer redisPool.Close()

	// Logger
	appLogger, accessLogger := initLogger()

	// Application Context
	appContext := &controllers.AppContext{
		DB:           db,
		Logger:       appLogger,
		AccessLogger: accessLogger,
		RedisPool:    redisPool,
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", viper.GetString("HTTP.HOST"), viper.GetString("HTTP.PORT")),
		Handler: router.New(appContext),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalln("Unable to start server", err.Error())
		}
	}()
	log.Printf("Start server at %s\n", srv.Addr)

	// Listen of Interrupt signals
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	// Wait for timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30)
	defer cancel()

	// Shutdown gracefully
	srv.Shutdown(ctx)
}

func initDB() *models.DB {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@%s/%s?parseTime=true", viper.GetString("MYSQL.USER"), viper.GetString("MYSQL.PASSWORD"), viper.GetString("MYSQL.HOST"), viper.GetString("MYSQL.DATABASE")))
	if err != nil {
		log.Fatalln("Unable to connect to database", err.Error())
	}

	return models.Use(db)
}

func initRedis() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   viper.GetInt("REDIS.MAX_IDLE"),
		MaxActive: viper.GetInt("REDIS.MAX_ACTIVE"),
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", viper.GetString("REDIS.HOST"), viper.GetString("REDIS.PORT")))
			if err != nil {
				log.Fatalln("Redis connection failed")
			}
			return c, err
		},
	}
}

func initLogger() (*logrus.Logger, *logrus.Logger) {
	appLogger := &logrus.Logger{
		Out: &lumberjack.Logger{
			Filename: "/home/praveen/go/src/github.com/praveen001/go-boilerplate/application.log",
			MaxSize:  5,
			MaxAge:   10,
			Compress: true,
		},
		Formatter:    &logrus.JSONFormatter{PrettyPrint: true},
		ReportCaller: true,
		Level:        logrus.InfoLevel,
	}

	accessLogger := &logrus.Logger{
		Out: &lumberjack.Logger{
			Filename: "/home/praveen/go/src/github.com/praveen001/go-boilerplate/access.log",
			MaxSize:  5,
			MaxAge:   10,
			Compress: true,
		},
		Level:     logrus.InfoLevel,
		Formatter: &logrus.TextFormatter{},
	}

	return appLogger, accessLogger
}
