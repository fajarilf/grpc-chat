package services

import (
	"context"

	pb "github.com/fajarilf/grpc-chat-proto/proto"
	"github.com/fajarilf/grpc-chat-server/models"
	"github.com/fajarilf/grpc-chat-server/repositories"
)

type RoomService struct {
	repo *repositories.RoomRepository
}

func NewRoomService(repo *repositories.RoomRepository) *RoomService {
	return &RoomService{
		repo: repo,
	}
}

func (rs *RoomService) CreateRoom(ctx context.Context, param *pb.RoomCreateRequest) *pb.RoomResponse {
	return &pb.RoomResponse{}
}

func (rs *RoomService) GetListRoom(ctx context.Context, param *pb.RoomListRequest) *pb.RoomListResponse {
	result, _ := rs.repo.Get(ctx, param)

	rooms := []*pb.RoomResponse{}
	for _, val := range result {
		room := models.ToRoomResponse(val)
		rooms = append(rooms, room)
	}

	return &pb.RoomListResponse{
		Rooms: rooms,
	}
}
