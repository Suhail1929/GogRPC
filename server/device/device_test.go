package device_test

import (
	"Bakri-Souhail/server/device"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Test_StoreToDatabase_CountOperations(t *testing.T) {
	deviceInstance := device.Device{
		DeviceName:       "DeviceTest",
		TotalOperations:  100,
		FailedOperations: 10,
	}

	err := deviceInstance.StoreToDatabase()
	if err != nil {
		t.Errorf("Failed to store device in database: %v", err)
	}
	
	//check if the device is stored in the database
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:root@10.11.2.97:27017"))
	if err != nil {
		t.Fatalf("failed to connect to MongoDB test server: %v", err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database("CompanyInfos").Collection("device")

	var storedDevice device.Device
	err = collection.FindOne(ctx, deviceInstance).Decode(&storedDevice)
	if err != nil {
		t.Fatalf("failed to retrieve device from database: %v", err)
	}
	
	// Assert that the retrieved device matches the original device
	assert.Equal(t, deviceInstance.DeviceName, storedDevice.DeviceName, "Device names should match")
	
	/*
	Test CountOperations
	*/

	// Cas de test avec un JSON représentant deux appareils avec différentes opérations
	jsonData := `[
		{
			"device_name": "Device1",
			"operations": [
				{"type": "operation1", "has_succeeded": true},
				{"type": "operation2", "has_succeeded": false},
				{"type": "operation3", "has_succeeded": true}
			]
		},
		{
			"device_name": "Device2",
			"operations": [
				{"type": "operation1", "has_succeeded": true},
				{"type": "operation2", "has_succeeded": true}
			]
		}
	]`

	expectedResult := map[string]map[string]int{
		"Device1": {
			"total_operations":  3,
			"failed_operations": 1,
		},
		"Device2": {
			"total_operations":  2,
			"failed_operations": 0,
		},
	}

	// Appel de la fonction countOperations avec le JSON de test
	result, err := device.CountOperations(jsonData)

	// Vérification des résultats
	assert.Nil(t, err, "erreur non attendue")
	assert.Equal(t, expectedResult, result, "résultat incorrect")


}
