package services

import (
	"context"

	pb "github.com/fajarilf/grpc-chat-proto/proto"
	"github.com/fajarilf/grpc-chat-server/models"
	"github.com/fajarilf/grpc-chat-server/repositories"
	"github.com/fajarilf/grpc-chat-server/utils"
	"google.golang.org/grpc/codes"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (us *UserService) Register(param *pb.UserCreateRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{}, nil
}

func (us *UserService) Login(param *pb.UserLoginRequest) (*pb.UserLoginResponse, error) {
	return &pb.UserLoginResponse{}, nil
}

func (us *UserService) GetList(ctx context.Context) (*pb.UserListResponse, error) {
	result, err := us.repo.Get(ctx)
	if err != nil {
		return nil, utils.WrapGRPCError(codes.Internal, "User Get List Error", err)
	}

	responses := make([]*pb.UserResponse, 0, len(result))
	for _, val := range result {
		responses = append(responses, models.ToUserResponse(val))
	}

	return &pb.UserListResponse{
		Users: responses,
	}, nil
}
