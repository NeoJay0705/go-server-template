package main

import (
	"context"
	"log"
	"time"

	pb "github.com/NeoJay0705/go-server-template/cmd/grpc/api"

	"google.golang.org/grpc"
)

func main() {
	// Connect to the server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a client
	client := pb.NewExampleServiceClient(conn)

	// Call the SayHello RPC
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.SayHello(ctx, &pb.HelloRequest{Name: "World"})
	if err != nil {
		log.Fatalf("Error calling SayHello: %v", err)
	}

	log.Printf("Response from server: %s", response.Message)
}
