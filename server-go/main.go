package main

import (
	"log"
	"net"

	pb "github.com/fajarilf/grpc-chat-server/proto"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	chatServer := NewChatServer()
	pb.RegisterChatServiceServer(grpcServer, chatServer)

	log.Println("Chat server started on :50051")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
