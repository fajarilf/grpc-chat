package services

import (
	pb "github.com/fajarilf/grpc-chat-proto/proto"
)

type RoomService struct{}

func NewRoomService() *RoomService {
	return &RoomService{}
}

func (rs *RoomService) CreateRoom(param *pb.RoomCreateRequest) *pb.RoomResponse {
	return &pb.RoomResponse{}
}

// GetListRoom returns a page of mock rooms. cursor is the index to start from
// and size caps the page (defaulting to the full list when size <= 0).
func (rs *RoomService) GetListRoom(param *pb.RoomListRequest) *pb.RoomListResponse {
	start := max(int(param.GetCursor()), 0)
	start = min(start, len(mockRooms))

	end := len(mockRooms)
	if size := int(param.GetSize()); size > 0 && start+size < end {
		end = start + size
	}

	return &pb.RoomListResponse{
		Rooms: mockRooms[start:end],
	}
}
