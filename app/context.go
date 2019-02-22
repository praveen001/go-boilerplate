package app

import (
	"context"
	"fmt"
	"net/http"

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

// New application with given configs
func New(conf *Config) *Context {
	c := &Context{
		Config: conf,
	}

	c.initLogger()
	c.initDB()
	c.initRedis()

	c.srv = &http.Server{
		Addr: fmt.Sprintf("%s:%s", conf.HTTP.Host, conf.HTTP.Port),
	}

	return c
}

// StartWith ..
func (c *Context) StartWith(router http.Handler) {
	c.srv.Handler = router

	go func() {
		if err := c.srv.ListenAndServe(); err != nil {
			c.Logger.Fatal("Unable to start server", err.Error())
		}
	}()
	c.Logger.Info("Start server at ", c.srv.Addr)
}

// Shutdown ..
func (c *Context) Shutdown(ctx context.Context) {
	c.DB.close()
	c.RedisPool.Close()

	c.srv.Shutdown(ctx)
}
