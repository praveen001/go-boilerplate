package controllers

import (
	"errors"
	"net/http"

	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/handlers"
	"github.com/praveen001/go-boilerplate/models"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

// AppContext holds the context for each request
// Everything in context must be thread-safe
type AppContext struct {
	DB        *models.DB
	RedisPool *redis.Pool
	Logger    *logrus.Logger
}

// RecoveryHandler returns 500 status when handler panics.
// Writes error to application log
func (c *AppContext) RecoveryHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		defer func() {
			if e := recover(); e != nil {
				switch t := e.(type) {
				case string:
					err = errors.New(t)
				case error:
					err = t
				default:
					err = errors.New("Unknown error")
				}
				c.Logger.Errorln(err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}()

		h.ServeHTTP(w, r)
	})
}

// CORSHandler handles cors requests
func (c *AppContext) CORSHandler(h http.Handler) http.Handler {
	return cors.New(cors.Options{
		AllowedHeaders: []string{"authorization", "content-type"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}).Handler(h)
}

// LogHandler writes access log
func (c *AppContext) LogHandler(h http.Handler) http.Handler {
	out := &lumberjack.Logger{
		Filename: "/home/praveen/go/src/github.com/praveen001/go-boilerplate/access.log",
		MaxSize:  10,
		MaxAge:   10,
		Compress: true,
	}
	return handlers.LoggingHandler(out, h)
}
