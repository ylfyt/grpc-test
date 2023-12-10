package main

import (
	"context"
	"fmt"
	"test_grpc/grpc/chat"
	"test_grpc/grpc/hello"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:5000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	helloClient := hello.NewHelloServiceClient(conn)
	res, err := helloClient.HelloWorld(context.Background(), &hello.HelloRequest{
		Body: "hello, world client",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Data: %+v\n", res)

	
	client := chat.NewChatServiceClient(conn)
	// Call the gRPC method
	response, err := client.SayHello(context.Background(), &chat.Message{
		Body: "Hello from client",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Data:", response.Body)
}
