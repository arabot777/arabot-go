package redis

import (
	"net"
	"strconv"
	"time"

	rdsv8 "github.com/go-redis/redis/v8"
)

type Config struct {
	Host     string
	Port     int
	Username string
	Passwd   string
	DB       int

	MaxRetries int

	PoolSize    int
	PoolTimeout time.Duration
	IdleTimeout time.Duration

	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func (cfg *Config) Options() *rdsv8.Options {
	opts := &rdsv8.Options{
		Network:  "tcp",
		Addr:     net.JoinHostPort(cfg.Host, strconv.Itoa(cfg.Port)),
		Username: cfg.Username,
		Password: cfg.Passwd,
		DB:       cfg.DB,

		MaxRetries:      cfg.MaxRetries,
		MinRetryBackoff: 8 * time.Millisecond,
		MaxRetryBackoff: 1024 * time.Millisecond,

		DialTimeout:  cfg.DialTimeout,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,

		PoolSize: cfg.PoolSize,
		// Amount of time client waits for connection if all connections
		// are busy before returning an error.
		// Default is ReadTimeout + 1 second.
		PoolTimeout: cfg.PoolTimeout,
		IdleTimeout: cfg.IdleTimeout,
	}
	if opts.MaxRetries > 3 {
		opts.MaxRetries = 3
	} else if opts.MaxRetries <= 0 {
		opts.MaxRetries = -1
	}
	if opts.DialTimeout > 3*time.Second || opts.DialTimeout < 50*time.Millisecond {
		opts.DialTimeout = 3 * time.Second
	}
	if opts.ReadTimeout > 3*time.Second || opts.ReadTimeout < 50*time.Millisecond {
		opts.ReadTimeout = 3 * time.Second
	}
	if opts.WriteTimeout > 3*time.Second || opts.WriteTimeout < 50*time.Millisecond {
		opts.WriteTimeout = 3 * time.Second
	}
	if opts.PoolSize > 20 {
		opts.PoolSize = 20
	}
	if opts.PoolSize < 3 {
		opts.PoolSize = 3
	}
	return opts
}
