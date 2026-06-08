package services

import (
	"math/rand/v2"
	"strconv"

	pb "github.com/fajarilf/grpc-chat-proto/proto"
)

type color_obj struct {
	label string
	value string
}

type img_obj struct {
	label string
	value string
}

// mockUsers are stand-in members referenced by the mock rooms below.
var mockUsers = []*pb.UserResponse{
	{Id: 1, Name: "Aiko Tanaka", Username: "aiko"},
	{Id: 2, Name: "Budi Santoso", Username: "budi"},
	{Id: 3, Name: "Chen Wei", Username: "chenw"},
	{Id: 4, Name: "Dewi Lestari", Username: "dewi"},
	{Id: 5, Name: "Eko Prabowo", Username: "eko"},
	{Id: 6, Name: "Fajar Ilham", Username: "fajar"},
}

var colors = []color_obj{
	{label: "Slate", value: "#475569"},
	{label: "Sky", value: "#0284c7"},
	{label: "Emerald", value: "#059669"},
	{label: "Amber", value: "#d97706"},
	{label: "Rose", value: "#e11d48"},
	{label: "Violet", value: "#7c3aed"},
	{label: "Ink", value: "#14110d"},
	{label: "Teal", value: "#0d9488"},
}

var imgs = []img_obj{
	{label: "Mountains", value: "https://picsum.photos/seed/mountains/640/480"},
	{label: "City", value: "https://picsum.photos/seed/city/640/480"},
	{label: "Ocean", value: "https://picsum.photos/seed/ocean/640/480"},
	{label: "Forest", value: "https://picsum.photos/seed/forest/640/480"},
	{label: "Desert", value: "https://picsum.photos/seed/desert/640/480"},
	{label: "Aurora", value: "https://picsum.photos/seed/aurora/640/480"},
}

func getRandomUsers(slice []*pb.UserResponse, count int) []*pb.UserResponse {
	if count > len(slice) {
		count = len(slice)
	}

	indices := rand.Perm(len(slice))
	result := make([]*pb.UserResponse, count)
	for i := range count {
		result[i] = slice[indices[i]]
	}

	return result
}

func generateMock(size int) []*pb.RoomResponse {
	rooms := make([]*pb.RoomResponse, 0, size)

	var (
		bg           string
		bg_type      string
		randNumColor int
		randNumImg   int
	)

	for i := range size {
		randNumColor = rand.IntN(len(colors))
		randNumImg = rand.IntN(len(imgs))

		// odd iteration get image background
		bg = imgs[randNumImg].value
		bg_type = "IMAGE"

		// even iteration get color background
		if i%2 == 0 {
			bg = colors[randNumColor].value
			bg_type = "COLOR"
		}

		users := getRandomUsers(mockUsers, 3)

		rooms = append(rooms, &pb.RoomResponse{
			Name:           "Room " + strconv.Itoa(i),
			Description:    "Mock Room " + strconv.Itoa(i),
			Background:     bg,
			BackgroundType: bg_type,
			Users:          users,
		})
	}

	return rooms
}

var mockRooms = generateMock(50)
