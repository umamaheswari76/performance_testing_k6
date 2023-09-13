package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	//"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const URI = "mongodb://localhost:27017/?retryWrites=true&connectTimeoutMS=10000"

var mongodb *mongo.Collection

type DbTest struct {
	Token string `json: "token"`
}

func main() {
	apiClient := API()
	mongodb = DBConnection()

	//setup routes
	apiClient.POST("/db_test", handleinsert())
	// apiClient.POST("/db_test_many",handlemany())
	apiClient.Run(":8000")
}

// initializing gin router
func API() *gin.Engine {
	r := gin.Default()
	return r
}

// mongodb connection
func DBConnection() *mongo.Collection {
	connection := options.Client().ApplyURI(URI)
	client, err := mongo.Connect(context.TODO(), connection)
	if err != nil {
		log.Fatal("Database is not connected", err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	fmt.Println("Database Connected")
	err1 := client.Ping(ctx, nil)
	if err1 != nil {
		log.Fatal("Mongodb is not available")
	}
	collection := client.Database("k6").Collection("db_test")
	fmt.Println("Collection Created")
	return collection
}

// insert a token endpoint
func handleinsert() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var temp DbTest
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&temp); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newTemp := DbTest{
			Token: temp.Token,
		}
		_, err1 := mongodb.InsertOne(ctx, newTemp)
		if err1 != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err1.Error()})
			return
		}
		c.JSON(http.StatusOK, newTemp)
		fmt.Println("Inserted succesfully")
	}
}

// func handlemany() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 		defer cancel()

// 		var tempslice = make([]DbTest, 0)
// 		lines := strings.Split(string(fileBytes), "\n")
// 		var temp DbTest

// 		//validating many
// 		if err := c.BindJSON(&tempslice); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{" error": err.Error()})
// 			return
// 		}

// 	}
// }
