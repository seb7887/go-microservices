package redis

import (
	"os"
	"os/signal"
	"syscall"
	"time"
	"github.com/gomodule/redigo/redis"
	"github.com/seb7887/go-microservices/config"
)

var (
	Pool *redis.Pool
)

func init() {
	redisHost := config.GetConfig().RedisHost
	Pool = newPool(redisHost)
	cleanupHook()
}

func newPool(server string) *redis.Pool {
	return &redis.Pool{
		MaxIdle: 3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func cleanupHook() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGKILL)
	go func() {
		<-c
		Pool.Close()
		os.Exit(0)
	}()
}