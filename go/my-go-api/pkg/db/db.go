// db.go
package db

import (
	"context"
	"fmt"
	"log"
	"my-go-api/pkg/config"
	"net/url"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func ConnectDB() *mongo.Client {
	// Load env variables from .env
	config.LoadConfig()
	// Fetch mongodb credentials from env variables
	username := url.QueryEscape(os.Getenv("MONGODB_USERNAME"))
	password := url.QueryEscape(os.Getenv("MONGODB_PASSWORD"))

	if username == "" || password == "" {
		log.Fatalf("MongoDB credentials are not set in environment variables")
	}

	uri := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.a8hjh.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0", username, password)

	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	var err error
	// Create a new Client and connect to the server
	client, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
		panic(err)
	}

	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		log.Fatalf("error pinging mongodb: %v", err)
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return client
}

func DisconnectDB() {
	if err := client.Disconnect(context.TODO()); err != nil {
		log.Fatalf("Error disconnecting from MongoDB: %v", err)
		panic(err)
	}
}
