package device

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"encoding/json"
)

type Device struct {
	DeviceName       string `json:"device_name"`
	TotalOperations  int    `json:"total_operations"`
	FailedOperations int    `json:"failed_operations"`
}

func (d *Device) StoreToDatabase() error {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://root:root@10.11.2.97:27017"))
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(context.Background())

	collection := client.Database("CompanyInfos").Collection("device")
	_, err = collection.InsertOne(context.Background(), d)
	if err != nil {
		return fmt.Errorf("failed to insert device into MongoDB: %v", err)
	}

	fmt.Printf("Device %s stored in database\n", d.DeviceName)

	return nil
}

func (d *Device) CountOperations(jsonData string) (map[string]map[string]int, error) {
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

