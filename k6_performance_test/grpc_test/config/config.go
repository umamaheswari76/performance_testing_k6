package config

import (
	"context"
	"fmt"
	"k6_performance_test/grpc_k6_test/constants"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDatabase()(*mongo.Collection){
	connection := options.Client().ApplyURI(constants.URI)
	client,err := mongo.Connect(context.TODO(), connection)

	if err != nil{
		log.Fatal("Database is not connected ",err)
	}
	fmt.Println("Database connected")
	
	ctx,_ := context.WithTimeout(context.Background(), 10*time.Second)
	err1 := client.Ping(ctx, nil)
	if err1 != nil{
		log.Fatal("Mongodb is not available")
	}

	collection := client.Database("k6").Collection("grpc_test")
	fmt.Println("Collection Created")
	return collection
}