package database

import ("context"
"log"
"time"
"fmt"
"go.mongodb.org/mongo-driver/mongo"
"go.mongodb.org/mongo-driver/mongo/options"
"go.mongodb.org/mongo-driver/mongo/readpref")


var Client *mongo.Client

func Connect() {
	// Set a timeout for the connection attempt
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	Client = client

	// Check connection with a timeout
	err = Client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("Could not connect to MongoDB:", err)
	} else {
		log.Println("Connected to MongoDB!")
	}
}

func GetCollection(collectionName string) *mongo.Collection {
	// Use context with timeout for MongoDB queries
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	fmt.Println("get collection called")
	// Return the MongoDB collection with the context
	return Client.Database("orders").Collection(collectionName)
	
}


