package main

import (
	"TikTokLite_v2/user_follow/user/pb"
	"TikTokLite_v2/user_follow/user/service"
	"context"
)

type Service struct {
	sv *service.UserService
}

func (s *Service) UserLogin(ctx context.Context, req *pb.UserLoginOrRegisterRequest) (*pb.UserLoginOrRegisterResponse, error) {
	resp, err := s.sv.UserLogin(ctx, req)
	return resp, err
}
func (s *Service) UserRegister(ctx context.Context, req *pb.UserLoginOrRegisterRequest) (*pb.UserLoginOrRegisterResponse, error) {
	resp, err := s.sv.UserRegister(ctx, req)
	return resp, err
}
func (s *Service) UserInfo(ctx context.Context, req *pb.UserInfoRequest) (*pb.UserInfoResponse, error) {
	resp, err := s.sv.UserInfo(ctx, req)
	return resp, err
}
