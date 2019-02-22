package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"go.uber.org/zap"

	"github.com/gomodule/redigo/redis"

	"github.com/praveen001/go-boilerplate/app"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

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
	appLogger := initLogger()

	// Application Context
	appContext := &app.Context{
		DB:        db,
		Logger:    appLogger,
		RedisPool: redisPool,
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

func initDB() *app.DB {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@%s/%s?parseTime=true", viper.GetString("MYSQL.USER"), viper.GetString("MYSQL.PASSWORD"), viper.GetString("MYSQL.HOST"), viper.GetString("MYSQL.DATABASE")))
	if err != nil {
		log.Fatalln("Unable to connect to database", err.Error())
	}

	return app.Use(db)
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

func initLogger() *app.Logger {
	config := zap.NewDevelopmentConfig()
	config.OutputPaths = []string{
		"/home/praveen/go/src/github.com/praveen001/go-boilerplate/application.log",
	}
	logger, _ := config.Build()

	return &app.Logger{logger.Sugar()}
}
