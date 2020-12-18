package service

import (
	"context"
	pb "demo/api/user/v1"
	"demo/internal/biz"
	"github.com/google/wire"
)

var Provider = wire.NewSet(New, wire.Bind(new(pb.UserServer), new(*Service)))

type Service struct {
	pb.UnimplementedUserServer
	uuc *biz.UserUseCase
}

func New(uuc *biz.UserUseCase) *Service {
	return &Service{uuc: uuc}
}

func (svr *Service) Login(ctx context.Context, r *pb.LoginRequest) (*pb.LoginReply, error) {
	// DTO -> DO
	o := new(biz.User)
	o.Account = r.Account

	svr.uuc.Login(o)

	return &pb.LoginReply{Message: "ok"}, nil
}