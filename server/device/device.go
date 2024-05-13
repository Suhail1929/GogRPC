package device

import (
	"context"
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client // Déclaration du client MongoDB comme variable globale

func init() {
	// Initialisation du client MongoDB au démarrage de l'application
	var err error
	client, err = mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://root:root@10.11.2.97:27017"))
	if err != nil {
		panic(fmt.Sprintf("failed to connect to MongoDB: %v", err))
	}
}

// Structure de données pour représenter un appareil
type Device struct {
	DeviceName       string `json:"device_name"`
	TotalOperations  int    `json:"total_operations"`
	FailedOperations int    `json:"failed_operations"`
}

// Méthode pour enregistrer les statistiques de l'appareil dans la collection 'device_statistics'
func (d *Device) StoreStatistics() error {
	collection := client.Database("CompanyInfos").Collection("device_statistics")
	_, err := collection.InsertOne(context.Background(), map[string]interface{}{
		"device_name":       d.DeviceName,
		"total_operations":  d.TotalOperations,
		"failed_operations": d.FailedOperations,
	})
	if err != nil {
		return fmt.Errorf("failed to insert device statistics into MongoDB: %v", err)
	}

	fmt.Printf("Device statistics for %s stored in 'device_statistics' collection\n", d.DeviceName)

	return nil
}

// Méthode pour enregistrer les détails des opérations de l'appareil dans la collection 'device_operations'
func (d *Device) StoreOperations(operations []struct {
	Type         string `json:"type"`
	HasSucceeded bool   `json:"has_succeeded"`
}) error {
	collection := client.Database("CompanyInfos").Collection("device_operations")
	_, err := collection.InsertOne(context.Background(), map[string]interface{}{
		"device_name": d.DeviceName,
		"operations":  operations,
	})
	if err != nil {
		return fmt.Errorf("failed to insert device operations into MongoDB: %v", err)
	}

	fmt.Printf("Operations for device %s stored in 'device_operations' collection\n", d.DeviceName)

	return nil
}

// Fonction pour compter les opérations, calculer les statistiques et enregistrer dans deux collections différentes
func CountOperations(jsonData string) (map[string]map[string]int, error) {
	var devicesData []struct {
		DeviceName string `json:"device_name"`
		Operations []struct {
			Type         string `json:"type"`
			HasSucceeded bool   `json:"has_succeeded"`
		} `json:"operations"`
	}

	if err := json.Unmarshal([]byte(jsonData), &devicesData); err != nil {
		return nil, fmt.Errorf("error decoding JSON: %w", err)
	}

	operationStats := make(map[string]map[string]int)

	// Loop through device data
	for _, devData := range devicesData {
		deviceName := devData.DeviceName
		totalOperations := len(devData.Operations)
		failedOperations := 0

		// Count operations and failed operations
		for _, op := range devData.Operations {
			if !op.HasSucceeded {
				failedOperations++
			}
		}

		// Create a Device instance to store statistics
		deviceInstance := &Device{
			DeviceName:       deviceName,
			TotalOperations:  totalOperations,
			FailedOperations: failedOperations,
		}

		// Store statistics in 'device_statistics' collection
		if err := deviceInstance.StoreStatistics(); err != nil {
			return nil, err
		}

		// Store operation details in 'device_operations' collection
		if err := deviceInstance.StoreOperations(devData.Operations); err != nil {
			return nil, err
		}

		// Add statistics to result map
		deviceStats := map[string]int{
			"total_operations":  totalOperations,
			"failed_operations": failedOperations,
		}
		operationStats[deviceName] = deviceStats
	}

	return operationStats, nil
}
