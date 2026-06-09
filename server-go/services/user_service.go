package services

import (
	pb "github.com/fajarilf/grpc-chat-proto/proto"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) Register(param *pb.UserCreateRequest) *pb.UserResponse {
	return &pb.UserResponse{}
}

func (us *UserService) Login(param *pb.UserLoginRequest) *pb.UserLoginResponse {
	return &pb.UserLoginResponse{}
}
