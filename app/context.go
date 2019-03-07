package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/praveen001/go-boilerplate/models"

	"github.com/gomodule/redigo/redis"
)

// Context ..
type Context struct {
	DB        *DB
	RedisPool *redis.Pool
	Logger    *Logger
	Config    *Config

	srv *http.Server
}

// New creates a new application with given configs
//
// It also initiates all application dependencies like DB connections
func New(conf *Config) *Context {
	c := &Context{
		Config: conf,
	}

	c.initLogger()
	c.initDB()
	c.initRedis()

	c.DB.AutoMigrate(models.Playlist{}, models.Item{})

	return c
}

// StartWith uses the given `http.Handler` for mapping HTTP requests
//
// Starts server in a goroutine
func (c *Context) StartWith(router http.Handler) {
	c.srv = &http.Server{
		Addr:    fmt.Sprintf("%s:%s", c.Config.HTTP.Host, c.Config.HTTP.Port),
		Handler: router,
	}

	go func() {
		if err := c.srv.ListenAndServe(); err != nil {
			c.Logger.Fatal("Unable to start server", err.Error())
		}
	}()
	c.Logger.Info("Start server at ", c.srv.Addr)
}

// Shutdown closes all the open connections, and finally shutsdown the HTTP server.
func (c *Context) Shutdown(ctx context.Context) {
	c.DB.Close()
	c.RedisPool.Close()

	c.srv.Shutdown(ctx)
}
