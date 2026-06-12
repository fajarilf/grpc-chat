package services

import (
	"context"
	"fmt"

	pb "github.com/fajarilf/grpc-chat-proto/proto"
	"github.com/fajarilf/grpc-chat-server/models"
	"github.com/fajarilf/grpc-chat-server/repositories"
)

type RoomService struct {
	repo         *repositories.RoomRepository
	userRoomRepo *repositories.UserRoomRepository
}

func NewRoomService(repo *repositories.RoomRepository, userRoomRepo *repositories.UserRoomRepository) *RoomService {
	return &RoomService{
		repo:         repo,
		userRoomRepo: userRoomRepo,
	}
}

func (rs *RoomService) CreateRoom(ctx context.Context, param *pb.RoomCreateRequest) (*pb.RoomResponseWithUser, error) {
	result, err := rs.repo.Create(ctx, param)
	if err != nil {
		return nil, fmt.Errorf("Create Room Error: %v", err)
	}

	userIDs := make([]int, len(param.UserIds))
	for i, id := range param.UserIds {
		userIDs[i] = int(id)
	}

	if err := rs.userRoomRepo.CreateBatch(ctx, result.Id, userIDs); err != nil {
		return nil, fmt.Errorf("Create Room Error: %v", err)
	}

	membersByRoom, err := rs.repo.MembersByRoomIDs(ctx, []int{result.Id})
	if err != nil {
		return nil, fmt.Errorf("Create Room Error: %v", err)
	}

	return models.ToRoomResponseWithUser(result, membersByRoom[result.Id]), nil
}

func (rs *RoomService) GetRoomById(ctx context.Context, param *pb.RoomId) (*pb.RoomResponseWithUser, error) {
	result, err := rs.repo.GetById(ctx, int(param.Id))
	if err != nil {
		return nil, fmt.Errorf("Get Room by Id Error: %v", err)
	}

	membersByRoom, err := rs.repo.MembersByRoomIDs(ctx, []int{result.Id})
	if err != nil {
		return nil, fmt.Errorf("Get Room by Id Error: %v", err)
	}

	return models.ToRoomResponseWithUser(result, membersByRoom[result.Id]), nil
}

func (rs *RoomService) GetListRoom(ctx context.Context, param *pb.RoomListRequest) (*pb.RoomListResponse, error) {
	result, err := rs.repo.Get(ctx, param)
	if err != nil {
		return nil, fmt.Errorf("Get List Room Error: %v", err)
	}

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
	}, nil
}
