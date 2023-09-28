package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
    // MongoDB connection URI
    uri := "mongodb://localhost:27017" 

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Set up MongoDB client 
    clientOptions := options.Client().ApplyURI(uri)

    // Connect to MongoDB
    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatal(err)
    }
    defer client.Disconnect(ctx)

    database := client.Database("mydb")

    // Name of the collection to be created (with a valid name)
    collectionName := "aadhaar_card_applications"

    // Create the collection
    err = database.CreateCollection(ctx, collectionName)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Collection '%s' created successfully.\n", collectionName)
}
