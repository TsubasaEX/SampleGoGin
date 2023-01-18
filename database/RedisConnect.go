package database

import (
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
)

var RedisDefaultPool *redis.Pool

func newPool(addr string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
		Dial: func() (redis.Conn, error) { return redis.Dial("tcp", addr) },
	}
}

func init() {
	log.Println("Initializing RedisConnect...")
	RedisDefaultPool = newPool("localhost:6379")
}
