package device

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Device struct {
	DeviceName       string `json:"device_name"`
	TotalOperations  int    `json:"total_operations"`
	FailedOperations int    `json:"failed_operations"`
}

func (d *Device) StoreToDatabase() error {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://root:root@localhost:27017"))
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

// package device

// import (
// 	"Souhail-Bakri/Go-gRPC/operations"
// 	"context"
// 	"fmt"

// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// type Device struct {
// 	device_name string                 `json: "device_name"`
// 	operations  []operations.Operation `json: "operations"`
// }

// func (d Device) PrettyPrint() error {
// 	fmt.Printf("Device %s has %d operations\n", d.device_name, len(d.operations))
// 	return nil
// }

// func (d Device) InitializeDevice(device_name string, operations []operations.Operation) (Device, error) {
// 	d.device_name = device_name
// 	d.operations = operations
// 	return d, nil
// }

// func (d Device) StoreToDatabase() error {

// 	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://root:root@localhost:27017"))
// 	coll := client.Database("CompanyInfos").Collection("device")
// 	coll.InsertOne(context.Background(), d, nil)

// 	if err != nil {
// 		return err
// 	}

// 	fmt.Printf("Device %s stored in database\n", d.device_name)

// 	return nil
// }
