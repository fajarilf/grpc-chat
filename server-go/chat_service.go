package main

import (
	"context"
	"fmt"
	"io"
	"sync"
	"time"

	pb "github.com/fajarilf/grpc-chat-server/proto"
)

type ChatServer struct {
	pb.UnimplementedChatServiceServer

	rooms     map[string]map[string]chan *pb.Message
	mu        sync.RWMutex
	history   map[string][]*pb.Message
	historyMu sync.RWMutex
}

func NewChatServer() *ChatServer {
	return &ChatServer{
		rooms:   make(map[string]map[string]chan *pb.Message),
		history: make(map[string][]*pb.Message),
	}
}

func (s *ChatServer) SendMessage(ctx context.Context, msg *pb.Message) (*pb.MessageResponse, error) {
	msg.Timestamp = time.Now().Unix()

	s.historyMu.Lock()
	s.history[msg.Room] = append(s.history[msg.Room], msg)
	s.historyMu.Unlock()

	s.mu.RLock()
	if clients, ok := s.rooms[msg.Room]; ok {
		for _, ch := range clients {
			select {
			case ch <- msg:
			default:
				//channel full, skip
			}
		}
	}
	s.mu.RUnlock()

	return &pb.MessageResponse{
		Success:   true,
		MessageId: fmt.Sprintf("%d", msg.Timestamp),
	}, nil
}

func (s *ChatServer) JoinChat(req *pb.JoinRequest, stream pb.ChatService_JoinChatServer) error {
	msgChan := make(chan *pb.Message, 100)

	s.mu.Lock()
	if s.rooms[req.Room] == nil {
		s.rooms[req.Room] = make(map[string]chan *pb.Message)
	}
	s.rooms[req.Room][req.Username] = msgChan
	s.mu.Unlock()

	joinMsg := &pb.Message{
		Username:  "System",
		Content:   fmt.Sprintf("%s joined the room", req.Username),
		Timestamp: time.Now().Unix(),
		Room:      req.Room,
	}
	s.SendMessage(context.Background(), joinMsg)

	defer func() {
		s.mu.Lock()
		delete(s.rooms[req.Room], req.Username)
		if len(s.rooms[req.Room]) == 0 {
			delete(s.rooms, req.Room)
		}
		s.mu.Unlock()
		close(msgChan)

		leaveMsg := &pb.Message{
			Username:  "System",
			Content:   fmt.Sprintf("%s left the room", req.Username),
			Timestamp: time.Now().Unix(),
			Room:      req.Room,
		}
		s.SendMessage(context.Background(), leaveMsg)
	}()

	for {
		select {
		case msg := <-msgChan:
			if err := stream.Send(msg); err != nil {
				return err
			}
		case <-stream.Context().Done():
			return nil
		}
	}
}

func (s *ChatServer) Chat(stream pb.ChatService_ChatServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		msg.Timestamp = time.Now().Unix()

		s.historyMu.Lock()
		s.history[msg.Room] = append(s.history[msg.Room], msg)
		s.historyMu.Unlock()

		s.mu.RLock()
		if clients, ok := s.rooms[msg.Room]; ok {
			for username, ch := range clients {
				if username != msg.Username {
					select {
					case ch <- msg:
					default:
					}
				}
			}
		}
		s.mu.RUnlock()

		if err := stream.Send(msg); err != nil {
			return err
		}
	}
}

func (s *ChatServer) GetHistory(ctx context.Context, req *pb.HistoryRequest) (*pb.HistoryResponse, error) {
	s.historyMu.RLock()
	defer s.historyMu.RUnlock()

	messages := s.history[req.Room]

	limit := int(req.Limit)
	if limit == 0 || limit > len(messages) {
		limit = len(messages)
	}

	start := max(len(messages)-limit, 0)

	return &pb.HistoryResponse{
		Message: messages[start:],
	}, nil
}
