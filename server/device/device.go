package device

import (
	"context"
	"fmt"
	"Souhail-Bakri/Go-gRPC/operations"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Device struct {
	device_name string                 `json: "device_name"`
	operations  []operations.Operation `json: "operations"`
}

func (d Device) PrettyPrint() error {
	fmt.Printf("Device %s has %d operations\n", d.device_name, len(d.operations))
	return nil
}

func (d Device) InitializeDevice(device_name string,operations []operations.Operation) (Device, error){
	d.device_name = device_name
	d.operations = operations
	return d, nil
}

func (d Device) StoreToDatabase() error {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://root:root@localhost:27017"))
	coll := client.Database("CompanyInfos").Collection("device")
	coll.InsertOne(context.Background(), d, nil)
	
	if err != nil {
		return err
	}

	fmt.Printf("Device %s stored in database\n", d.device_name)

	return nil
}


func countOperations() {
	
}


