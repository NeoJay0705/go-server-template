package main

import (
	"context"
	"log"
	"net"

	pb "github.com/NeoJay0705/go-server-template/cmd/grpc/api"

	"google.golang.org/grpc"
)

// Server struct
type server struct {
	pb.UnimplementedExampleServiceServer
}

// Implement the SayHello method
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello, " + req.Name}, nil
}

func main() {
	// Listen on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server
	s := grpc.NewServer()
	pb.RegisterExampleServiceServer(s, &server{})

	log.Println("Server is running on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
