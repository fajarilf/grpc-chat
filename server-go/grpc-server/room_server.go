package grpc_server

import (
	"context"

	pb "github.com/fajarilf/grpc-chat-proto/proto"
	"github.com/fajarilf/grpc-chat-server/repositories"
	"github.com/fajarilf/grpc-chat-server/services"
)

// RoomServer adapts the gRPC RoomServiceServer interface onto the plain
// services.RoomService, translating the context/error-bearing RPC signatures
// to the service layer's simpler method shapes.
type RoomServer struct {
	pb.UnimplementedRoomServiceServer

	service *services.RoomService
}

func NewRoomServer(repo *repositories.RoomRepository) *RoomServer {
	return &RoomServer{service: services.NewRoomService(repo)}
}

func (s *RoomServer) CreateRoom(ctx context.Context, req *pb.RoomCreateRequest) (*pb.RoomResponse, error) {
	return s.service.CreateRoom(ctx, req), nil
}

func (s *RoomServer) GetListRoom(ctx context.Context, req *pb.RoomListRequest) (*pb.RoomListResponse, error) {
	return s.service.GetListRoom(ctx, req), nil
}
