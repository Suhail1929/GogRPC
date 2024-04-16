package operations

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Operation struct {
	Type string `json: "type"`
	Has_succeeded bool `json: "has_succeeded"`
}

func (o Operation) InitializeOperation(Type string, Has_succeeded bool) (Operation, error){
	o.Type = Type
	o.Has_succeeded = Has_succeeded
	return o, nil
}


func (o Operation) StoreToDatabase() error{
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://root:root@localhost:27017"))
	coll := client.Database("CompanyInfos").Collection("operations")
	coll.InsertOne(context.Background(), o, nil)
	
	if err != nil {
		return err
	}

	fmt.Printf("Operation %s stored in database\n", o.Type)
	return nil
}