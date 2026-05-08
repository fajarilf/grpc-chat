package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	pb "github.com/fajarilf/grpc-chat-client/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewChatServiceClient(conn)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Println("Enter room name: ")
	room, _ := reader.ReadString('\n')
	room = strings.TrimSpace(room)

	stream, err := client.JoinChat(context.Background(), &pb.JoinRequest{
		Username: username,
		Room:     room,
	})
	if err != nil {
		log.Fatalf("Failed to join chat: %v", err)
	}

	go func() {
		for {
			msg, err := stream.Recv()
			if err != nil {
				log.Printf("Error receiving message: %v", err)
				return
			}

			timestamp := time.Unix(msg.Timestamp, 0).Format("15:04:05")
			fmt.Printf("\r\033[2K[%s] %s: %s\n>> ", timestamp, msg.Username, msg.Content)
		}
	}()

	fmt.Println("Connected! Type your message (or /quit to exit): ")
	fmt.Print(">> ")
	for {
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if text == "/quit" {
			fmt.Println("Goodbye!")
			return
		}

		if text == "" {
			continue
		}

		fmt.Print("\033[1A\033[2K")

		_, err := client.SendMessage(context.Background(), &pb.Message{
			Username: username,
			Content:  text,
			Room:     room,
		})
		if err != nil {
			log.Printf("Failed to send message: %v", err)
			fmt.Print("> ")
		}
	}
}
