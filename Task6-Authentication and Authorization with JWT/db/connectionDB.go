package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var TaskCollection *mongo.Collection
var UserCollection *mongo.Collection

func ConnectDB() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    mongoURI := os.Getenv("MONGODB_URL")   

   
    clientOptions := options.Client().ApplyURI(mongoURI)

    client, err := mongo.NewClient(clientOptions)

    if err != nil {
        log.Fatalf("Error creating MongoDB client: %v", err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err = client.Connect(ctx)
    if err != nil {
        log.Fatalf("Error connecting to MongoDB: %v", err)
    }

    Client = client
    TaskCollection = client.Database("jwt_task_manegnment").Collection("tasks")
    UserCollection = client.Database("jwt_task_manegnment").Collection("users")
}
