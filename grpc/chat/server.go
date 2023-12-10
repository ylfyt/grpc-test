package chat

import (
	"context"
	"fmt"
)

type ChatServer struct{
	UnimplementedChatServiceServer
}

func (s *ChatServer) SayHello(ctx context.Context, req *Message) (*Message, error) {
	fmt.Println("Data:", req.Body)
	return &Message{Body: "Hello from server"}, nil
}
