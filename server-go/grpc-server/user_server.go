package grpc_server

import (
	"context"

	pb "github.com/fajarilf/grpc-chat-proto/proto"
	"github.com/fajarilf/grpc-chat-server/services"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer

	service *services.UserService
}

func NewUserServer() *UserServer {
	return &UserServer{service: services.NewUserService()}
}

func (s *UserServer) Register(ctx context.Context, req *pb.UserCreateRequest) (*pb.UserResponse, error) {
	return s.service.Register(req), nil
}

func (s *UserServer) Login(ctx context.Context, req *pb.UserLoginRequest) (*pb.UserLoginResponse, error) {
	return s.service.Login(req), nil
}
