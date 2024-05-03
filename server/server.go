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
	//pour pouvoir utiliser les méthodes de la structure Device
	operationStats, err := device.CountOperations(req.GetMessage())
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


