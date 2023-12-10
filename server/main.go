package main

import (
	"fmt"
	"net"
	"test_grpc/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		panic(err)
	}

	s := grpc.NewRpcServer()
	fmt.Println("Listening on port", 5000)
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
