package models

import (
	"github.com/astaxie/beego"
	"github.com/gomodule/redigo/redis"
	"time"
)

func init() {
	redisServer := beego.AppConfig.String("redisServer")
	redisPassword := beego.AppConfig.String("redisPassword")
	redisDb := beego.AppConfig.String("redisDb")
	Pool = newPool(redisServer, redisPassword, redisDb)
}

var (
	Pool *redis.Pool
)

func newPool(server, password, db string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		MaxActive:   5,
		Wait:        true,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("AUTH", password); err != nil {
				c.Close()
				return nil, err
			}
			if _, err := c.Do("SELECT", db); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
	}
}
