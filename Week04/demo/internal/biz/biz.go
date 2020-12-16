package biz

import (
	"github.com/google/wire"
)

var Provider = wire.NewSet(New)

type User struct {
	Account string
}

type UserRepo interface {
	Auth(user *User)
}

func New(repo UserRepo) *UserUseCase {
	return &UserUseCase{repo: repo}
}

type UserUseCase struct {
	repo UserRepo
}

func (u *UserUseCase) Login(o *User) {
	u.repo.Auth(o)
}