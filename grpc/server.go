package grpc

import (
	"test_grpc/grpc/chat"
	"test_grpc/grpc/hello"

	"google.golang.org/grpc"
)

func NewRpcServer() *grpc.Server {
	s := grpc.NewServer()
	chat.RegisterChatServiceServer(s, &chat.ChatServer{})
	hello.RegisterHelloServiceServer(s, &hello.HelloServer{})

	return s
}
