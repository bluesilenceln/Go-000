package data

import (
	"database/sql"
	"demo/internal/biz"
	"demo/pkg/cache/redis"
	_ "demo/pkg/database/sql"
	mysql "demo/pkg/database/sql"
	"github.com/google/wire"
)

var Provider = wire.NewSet(New, wire.Bind(new(biz.UserRepo), new(*userRepo)), redis.NewRedis, mysql.NewMySQL)

// 编译时，可知道该对象实现了外部的某接口
var _ biz.UserRepo = (biz.UserRepo)(nil)

func New(db *sql.DB, r *redis.Redis) *userRepo {
	u := new(userRepo)
	u.db = db
	u.redis = r
	return u
}

type userRepo struct {
	db *sql.DB
	redis *redis.Redis
}

func (ur *userRepo) Auth(user *biz.User) {
	// select id where account = user.Account
	return
}
