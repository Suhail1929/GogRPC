package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"

	"Bakri-Souhail/GoGrpcServer/device"
	pb "Bakri-Souhail/GoGrpcServer/grpc"

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
	operationStats, err := countOperations(req.GetMessage())
	if err != nil {
		fmt.Println("Erreur lors du comptage des opérations :", err)
		return nil, err
	}

	for deviceName, stats := range operationStats {
		totalOperations := stats["total_operations"]
		failedOperations := stats["failed_operations"]
		fmt.Printf("Device Name: %s | Total Operations: %d | Failed Operations: %d\n", deviceName, totalOperations, failedOperations)

		deviceInstance := device.Device{
			DeviceName:       deviceName,
			TotalOperations:  totalOperations,
			FailedOperations: failedOperations,
		}

		err := deviceInstance.StoreToDatabase()
		if err != nil {
			fmt.Printf("Failed to store device %s in database: %v\n", deviceName, err)
		} else {
			fmt.Printf("Device %s stored in database\n", deviceName)
		}
	}

	return &pb.StringResponse{Message: "Server received: " + req.GetMessage()}, nil
}

func countOperations(jsonData string) (map[string]map[string]int, error) {
	var devicesData []struct {
		DeviceName string `json:"device_name"`
		Operations []struct {
			Type         string `json:"type"`
			HasSucceeded bool   `json:"has_succeeded"`
		} `json:"operations"`
	}

	if err := json.Unmarshal([]byte(jsonData), &devicesData); err != nil {
		return nil, fmt.Errorf("erreur lors du décodage JSON : %w", err)
	}

	operationStats := make(map[string]map[string]int)

	for _, devData := range devicesData {
		deviceName := devData.DeviceName
		totalOperations := len(devData.Operations)
		failedOperations := 0

		for _, op := range devData.Operations {
			if !op.HasSucceeded {
				failedOperations++
			}
		}

		deviceStats := map[string]int{
			"total_operations":  totalOperations,
			"failed_operations": failedOperations,
		}

		operationStats[deviceName] = deviceStats
	}

	return operationStats, nil
}

// func countOperations(jsonData string) (map[string]map[string]int, error) {
// 	// Décodez les données JSON dans une structure appropriée
// 	var devices []device.Device
// 	if err := json.Unmarshal([]byte(jsonData), &devices); err != nil {
// 		return nil, fmt.Errorf("erreur lors du décodage JSON : %w", err)
// 	}

// 	// Créez une carte pour stocker les statistiques d'opérations pour chaque device
// 	operationStats := make(map[string]map[string]int)

// 	// Parcourez chaque device pour compter les opérations
// 	for _, device := range devices {
// 		totalOperations := len(device.Operations)
// 		failedOperations := 0

// 		// Parcourez les opérations pour compter le nombre d'opérations échouées
// 		for _, op := range device.Operations {
// 			if !op.Has_succeeded {
// 				failedOperations++
// 			}
// 		}

// 		// Stockez les statistiques d'opérations pour ce device
// 		deviceStats := map[string]int{
// 			"total_operations":  totalOperations,
// 			"failed_operations": failedOperations,
// 		}

// 		operationStats[device.DeviceName] = deviceStats
// 	}

// 	return operationStats, nil
// }

// func countOperations(jsonData string) (map[string]map[string]int, error) {
// 	var devices []device.Device
// 	if err := json.Unmarshal([]byte(jsonData), &devices); err != nil {
// 		return nil, fmt.Errorf("erreur lors du décodage JSON : %w", err)
// 	}

// 	operationStats := make(map[string]map[string]int)

// 	// Parcourez chaque device pour compter les opérations
// 	for _, device := range devices {
// 		totalOperations := len(device.Operations)
// 		failedOperations := 0

// 		// Parcourez les opérations pour compter le nombre d'opérations échouées
// 		for _, op := range device.Operations {
// 			if !op.Has_succeeded {
// 				failedOperations++
// 			}
// 		}

// 		// Stockez les statistiques d'opérations pour ce device
// 		deviceStats := map[string]int{
// 			"total_operations":  totalOperations,
// 			"failed_operations": failedOperations,
// 		}

// 		operationStats[device.DeviceName] = deviceStats
// 	}

// 	for _, device := range devices {
// 		totalOperations := device.TotalOperations
// 		failedOperations := device.FailedOperations

// 		operationStats[device.DeviceName] = map[string]int{
// 			"total_operations":  totalOperations,
// 			"failed_operations": failedOperations,
// 		}
// 	}

// 	return operationStats, nil
// }
