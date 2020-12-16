package service

import (
	"context"
	v1 "demo/api/user/v1"
	"demo/internal/biz"
	"github.com/google/wire"
)

var Provider = wire.NewSet(New, wire.Bind(new(v1.UserServer), new(*Service)))

type Service struct {
	v1.UnimplementedUserServer
	uuc *biz.UserUseCase
}

func New(uuc *biz.UserUseCase) *Service {
	return &Service{uuc: uuc}
}

func (svr *Service) Login(ctx context.Context, r *v1.LoginRequest) (*v1.LoginReply, error) {
	// DTO -> DO
	o := new(biz.User)
	o.Account = r.Account

	svr.uuc.Login(o)

	return &v1.LoginReply{Message: "ok"}, nil
}