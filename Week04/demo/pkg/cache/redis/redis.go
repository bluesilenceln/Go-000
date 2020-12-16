package redis

import (
	"github.com/go-redis/redis/v8"
	"time"
)

type Config struct {
	Addrs []string
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Redis struct {
	cluster *redis.ClusterClient
}

func NewRedis() *Redis {
	r := new(Redis)
	//r.cluster = redis.NewClusterClient(&redis.ClusterOptions{
	//	Addrs: c.Addrs,
	//})
	return r
}
