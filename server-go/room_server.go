package main

import (
	"context"

	roompb "github.com/fajarilf/grpc-chat-proto/proto"
	"github.com/fajarilf/grpc-chat-server/services"
)

// RoomServer adapts the gRPC RoomServiceServer interface onto the plain
// services.RoomService, translating the context/error-bearing RPC signatures
// to the service layer's simpler method shapes.
type RoomServer struct {
	roompb.UnimplementedRoomServiceServer

	service *services.RoomService
}

func NewRoomServer() *RoomServer {
	return &RoomServer{service: services.NewRoomService()}
}

func (s *RoomServer) CreateRoom(ctx context.Context, req *roompb.RoomCreateRequest) (*roompb.RoomResponse, error) {
	return s.service.CreateRoom(req), nil
}

func (s *RoomServer) GetListRoom(ctx context.Context, req *roompb.RoomListRequest) (*roompb.RoomListResponse, error) {
	return s.service.GetListRoom(req), nil
}
