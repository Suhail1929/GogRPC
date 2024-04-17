package main

import (
	"context"
	"log"
	"net"

	pb "Bakri-Souhail/GoGrpcServer/grpc"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedStringServiceServer
}

func (s *server) SendString(ctx context.Context, req *pb.StringRequest) (*pb.StringResponse, error) {
	log.Printf("Received: %v", req.GetMessage())
	return &pb.StringResponse{Message: "Server received: " + req.GetMessage()}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterStringServiceServer(s, &server{})
	log.Println("Server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
