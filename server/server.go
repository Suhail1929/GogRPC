package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"

	pb "Bakri-Souhail/GoGrpcServer/grpc"
	"Bakri-Souhail/GoGrpcServer/operations"

	"google.golang.org/grpc"
)

// Device représente un appareil avec ses opérations
type Device struct {
	DeviceName string                 `json:"device_name"`
	Operations []operations.Operation `json:"operations"`
}

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
	// log.Printf("Received: %v", req.GetMessage())
	// Compter les opérations pour chaque device
	operationStats, err := countOperations(req.GetMessage())
	if err != nil {
		fmt.Println("Erreur lors du comptage des opérations :", err)
		log.Fatal(err)
	}

	// Afficher les statistiques d'opérations pour chaque device
	for deviceName, stats := range operationStats {
		totalOperations := stats["total_operations"]
		failedOperations := stats["failed_operations"]
		fmt.Printf("Device Name: %s | Total Operations: %d | Failed Operations: %d\n", deviceName, totalOperations, failedOperations)
	}
	return &pb.StringResponse{Message: "Server received: " + req.GetMessage()}, nil
}

func countOperations(jsonData string) (map[string]map[string]int, error) {
	// Décodez les données JSON dans une structure appropriée
	var devices []Device
	if err := json.Unmarshal([]byte(jsonData), &devices); err != nil {
		return nil, fmt.Errorf("erreur lors du décodage JSON : %w", err)
	}

	// Créez une carte pour stocker les statistiques d'opérations pour chaque device
	operationStats := make(map[string]map[string]int)

	// Parcourez chaque device pour compter les opérations
	for _, device := range devices {
		totalOperations := len(device.Operations)
		failedOperations := 0

		// Parcourez les opérations pour compter le nombre d'opérations échouées
		for _, op := range device.Operations {
			if !op.Has_succeeded {
				failedOperations++
			}
		}

		// Stockez les statistiques d'opérations pour ce device
		operationStats[device.DeviceName] = map[string]int{
			"total_operations":  totalOperations,
			"failed_operations": failedOperations,
		}
	}

	return operationStats, nil
}
