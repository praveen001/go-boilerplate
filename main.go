package main

import (
	"context"
	"flag"
	"os"
	"os/signal"

	"github.com/praveen001/go-boilerplate/app"
	"github.com/praveen001/go-boilerplate/router"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/spf13/viper"
)

func init() {
	configPath := flag.String("configPath", "./", "Config file path")
	flag.Parse()

	viper.AddConfigPath(*configPath)
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func main() {
	var conf app.Config

	if len(flag.Args()) == 0 {
		panic("Specify environment")
	}
	environ := flag.Args()[0]
	viper.UnmarshalKey(environ, &conf)
	conf.Environment = app.Environment(environ)

	app := app.New(&conf)
	app.StartWith(router.New(app))

	// Listen of Interrupt signals
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	// Wait for timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30)
	defer cancel()

	// Shutdown gracefully
	app.Shutdown(ctx)
}
