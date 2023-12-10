package hello

import (
	"context"
	"fmt"
)

type HelloServer struct {
	UnimplementedHelloServiceServer
}

func (s *HelloServer) HelloWorld(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	fmt.Println("Data:", req.Body)
	return &HelloResponse{Body: "Hello from server"}, nil
}
