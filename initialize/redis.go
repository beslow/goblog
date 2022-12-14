package initialize

import (
	"time"

	"github.com/beslow/goblog/config"
	"github.com/gomodule/redigo/redis"
)

var RedisPool *redis.Pool

func init() {
	RedisPool = &redis.Pool{
		MaxIdle:     config.Redis.MaxIdel,
		IdleTimeout: time.Duration(config.Redis.IdleTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial(
				"tcp",
				config.Redis.Host+":"+config.Redis.Port,
				redis.DialUsername(config.Redis.Username),
				redis.DialPassword(config.Redis.Password),
			)
		},
	}

	conn := RedisPool.Get()
	if _, err := conn.Do("ping"); err != nil {
		panic(err)
	}
}
