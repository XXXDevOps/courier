package models

import (
	"fmt"
	"github.com/NioDevOps/courier/cfg"
	"github.com/gomodule/redigo/redis"
	"time"
)

var (
	RedisPool *redis.Pool
)

func Init(c cfg.RedisCfgStruct) {
	redisHost := fmt.Sprintf("%s:%d", c.Host, c.Port)
	RedisPool = newRedisPool(redisHost, c.PoolSize, c.IdleTimeout)
}

func newRedisPool(addr string, pool_size int, idle_timeout time.Duration) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     pool_size,
		IdleTimeout: idle_timeout * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", addr) },
	}
}
