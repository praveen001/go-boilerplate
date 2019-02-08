package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/praveen001/go-boilerplate/models"

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
	// MySQL
	db := models.InitDB()
	defer db.Close()

	appContext := &controllers.AppContext{
		DB: db,
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", viper.GetString("HTTP.HOST"), viper.GetString("HTTP.PORT")),
		Handler: router.InitRouter(appContext),
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
	os.Exit(0)
}
