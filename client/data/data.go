package data

import (
	pb "Bakri-Souhail/GoGrpcClient/grpc"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"google.golang.org/grpc"
)

type Data struct {
	Json string
}

const (
	address = "10.11.2.97:50051"
)

func OpenJsonFile(choice string,path string) (*os.File, error) {
	filePath := path +"journee_" + choice + ".json"
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func ShowDevices(choice string) {
	file, err := OpenJsonFile(choice, "data/")
	if err != nil {
		log.Fatalf("Error opening JSON file: %v", err)
	}
	defer file.Close()

	// Read the content of the JSON file
	fileData, err := ReadFiles(file)
	if err != nil {
		log.Fatalf("Error reading JSON file: %v", err)
	}

	// Now you can access the JSON content as a string using fileData.Content
	fmt.Printf("File %s send to server\n", choice)
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
}

// ReadFiles reads the content of a JSON file and returns it as a Data object
func ReadFiles(file *os.File) (Data, error) {
	// Read the JSON file
	jsonData, err := ioutil.ReadAll(file)
	if err != nil {
		return Data{}, fmt.Errorf("failed to read JSON file: %w", err)
	}

	return Data{Json: string(jsonData)}, nil
}
