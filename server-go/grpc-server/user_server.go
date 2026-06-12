package grpc_server

import (
	"context"

	pb "github.com/fajarilf/grpc-chat-proto/proto"
	"github.com/fajarilf/grpc-chat-server/repositories"
	"github.com/fajarilf/grpc-chat-server/services"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer

	service *services.UserService
}

func NewUserServer(repo *repositories.UserRepository) *UserServer {
	return &UserServer{service: services.NewUserService(repo)}
}

func (s *UserServer) Register(ctx context.Context, req *pb.UserCreateRequest) (*pb.UserResponse, error) {
	return s.service.Register(req)
}

func (s *UserServer) Login(ctx context.Context, req *pb.UserLoginRequest) (*pb.UserLoginResponse, error) {
	return s.service.Login(req)
}

func (s *UserServer) GetList(ctx context.Context, empty *emptypb.Empty) (*pb.UserListResponse, error) {
	return s.service.GetList(ctx)
}
