package myRedis

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

var GlobalRedisPool *redis.Pool

func InitPool(address string, maxIdle, maxActie int, idleTimeout time.Duration) {
	GlobalRedisPool = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActie,
		IdleTimeout: idleTimeout,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", address)
		},
	}
}
