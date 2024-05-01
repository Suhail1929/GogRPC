package main

import (
	"Bakri-Souhail/GoGrpcClient/data"
	pb "Bakri-Souhail/GoGrpcClient/grpc"
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"
)

const (
	address = "10.11.2.97:50051"
)

func main() {
	ShowDevices()
}



func OpenJsonFile() (*os.File, error) {
	Choice := "1"
	filePath := "data/journee_" + Choice + ".json"
	// filePath := "data/journee_1.json"
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func ShowDevices() {
	file, err := OpenJsonFile()
	if err != nil {
		log.Fatalf("Error opening JSON file: %v", err)
	}
	defer file.Close()

	// Read the content of the JSON file
	fileData, err := data.ReadFiles(file)
	if err != nil {
		log.Fatalf("Error reading JSON file: %v", err)
	}

	// Now you can access the JSON content as a string using fileData.Content
	fmt.Println("Data send")
	SendToServer(fileData.Json)
}

func SendToServer(s string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewStringServiceClient(conn)

	resp, err := client.SendString(context.Background(), &pb.StringRequest{Message: s})
	if err != nil {
		log.Fatalf("Error calling SendString: %v", err)
	}

	log.Printf("Response from server: %s", resp.Message)

	os.Exit(0)
}
