package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const databaseName = "user-management-database"

func ConnectDB() *mongo.Client {

	mongoURI := os.Getenv("MONGO_DATABASE_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://mongodb:27017"
	}

	clientOptions := options.Client().ApplyURI(mongoURI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err1 := mongo.Connect(ctx, clientOptions)
	if err1 != nil {
		panic(err1)
	}

	err2 := client.Ping(ctx, nil)
	if err2 != nil {
		panic(err2)
	}

	log.Printf("Connected to MongoDB!")

	return client
}

var DB *mongo.Client = ConnectDB()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database(databaseName).Collection(collectionName)
	return collection
}

func CloseDB() {
	if DB == nil {
		return
	}

	err := DB.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}

	fmt.Println("Connection to MongoDB closed.")
}
