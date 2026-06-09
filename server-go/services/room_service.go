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

func (rs *RoomService) CreateRoom(ctx context.Context, param *pb.RoomCreateRequest) *pb.RoomResponseWithUser {
	result, _ := rs.repo.Create(ctx, param)

	membersByRoom, _ := rs.repo.MembersByRoomIDs(ctx, []int{result.Id})

	return models.ToRoomResponseWithUser(result, membersByRoom[result.Id])
}

func (rs *RoomService) GetRoomById(ctx context.Context, param *pb.RoomId) *pb.RoomResponseWithUser {
	result, _ := rs.repo.GetById(ctx, int(param.Id))

	membersByRoom, _ := rs.repo.MembersByRoomIDs(ctx, []int{result.Id})

	return models.ToRoomResponseWithUser(result, membersByRoom[result.Id])
}

func (rs *RoomService) GetListRoom(ctx context.Context, param *pb.RoomListRequest) *pb.RoomListResponse {
	result, _ := rs.repo.Get(ctx, param)

	ids := make([]int, len(result.Rooms))
	for i, val := range result.Rooms {
		ids[i] = val.Id
	}

	membersByRoom, _ := rs.repo.MembersByRoomIDs(ctx, ids)

	rooms := make([]*pb.RoomResponseWithUser, 0, len(result.Rooms))
	for _, val := range result.Rooms {
		room := models.ToRoomResponseWithUser(val, membersByRoom[val.Id])
		rooms = append(rooms, room)
	}

	return &pb.RoomListResponse{
		Rooms:      rooms,
		NextCursor: int32(result.NextCursor),
		PrevCursor: int32(result.PrevCursor),
		HasMore:    result.HasMore,
	}
}
