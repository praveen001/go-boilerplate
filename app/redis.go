package app

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func (c *Context) initRedis() {
	c.RedisPool = &redis.Pool{
		MaxIdle:   c.Config.Redis.MaxIdle,
		MaxActive: c.Config.Redis.MaxActive,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", c.Config.Redis.Host, c.Config.Redis.Port))
			if err != nil {
				c.Logger.Fatal("Redis connection failed", err.Error())
			}
			return conn, err
		},
	}
}
