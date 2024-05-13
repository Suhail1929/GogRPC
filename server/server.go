package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"Bakri-Souhail/server/device"
	pb "Bakri-Souhail/server/grpc"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedStringServiceServer
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

func (s *server) SendString(ctx context.Context, req *pb.StringRequest) (*pb.StringResponse, error) {
	// Appel de la fonction CountOperations du package device
	operationStats, err := device.CountOperations(req.GetMessage())
	if err != nil {
		fmt.Println("Erreur lors du comptage des opérations :", err)
		return nil, err
	}

	// Manipuler les résultats des opérations ici (par exemple, afficher ou utiliser les statistiques)

	// Construire une réponse à renvoyer au client
	responseMessage := fmt.Sprintf("Server received: %s\n", req.GetMessage())
	for deviceName, stats := range operationStats {
		responseMessage += fmt.Sprintf("Device: %s\n", deviceName)
		responseMessage += fmt.Sprintf("Total Operations: %d\n", stats["total_operations"])
		responseMessage += fmt.Sprintf("Failed Operations: %d\n", stats["failed_operations"])
	}

	return &pb.StringResponse{Message: responseMessage}, nil
}
