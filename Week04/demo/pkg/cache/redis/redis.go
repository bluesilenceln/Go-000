package redis

import "demo/pkg/config/cache/redis"

type Redis struct {
	Config *redis.Config
}

func New(config *redis.Config) *Redis {
	r := new(Redis)
	r.Config = config

	return r
}

func (r *Redis) Close() {

}