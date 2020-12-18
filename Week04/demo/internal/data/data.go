package data

import (
	"demo/internal/biz"
	"demo/pkg/cache/redis"
	"github.com/google/wire"
)

var Provider = wire.NewSet(New, wire.Bind(new(biz.UserRepo), new(*userRepo)), NewRedis)

// 编译时，可知道该对象实现了外部的某接口
var _ biz.UserRepo = (biz.UserRepo)(nil)

func New(r *redis.Redis) *userRepo {
	u := new(userRepo)
	u.redis = r
	return u
}

type userRepo struct {
	redis *redis.Redis
}

func (ur *userRepo) Auth(user *biz.User) {
	// select id where account = user.Account
	return
}
