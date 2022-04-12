package redis

import (
	"context"
	"fmt"
	"github.com/arabot777/arabot-go/pkg/logger"
	"net"
	"strconv"
	"time"

	rdsv8 "github.com/go-redis/redis/v8"
)

const Nil = rdsv8.Nil

type Client struct {
	rdsv8.Client
}

func InitRedis(conf *Config) (*Client, error) {
	c := &Client{*rdsv8.NewClient(conf.Options())}
	err := c.Ping(context.Background()).Err()
	if err != nil {
		logger.Fatalf(fmt.Sprintf("redis: can't connect redis, address: %s", net.JoinHostPort(conf.Host, strconv.Itoa(conf.Port))), err)
		return nil, err
	}
	return c, err
}

func RedisFactory(host string, port int, username, password string, db int) (*Client, error) {
	return InitRedis(&Config{
		Host:         host,
		Port:         port,
		Username:     username,
		Passwd:       password,
		DB:           db,
		MaxRetries:   3,
		PoolSize:     10,
		PoolTimeout:  time.Second * 3,
		IdleTimeout:  time.Second * 60 * 5,
		DialTimeout:  time.Second,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	})
}
