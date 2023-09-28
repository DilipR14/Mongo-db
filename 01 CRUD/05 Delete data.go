package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type Applicant struct {
    UAN         string `bson:"uan"`
    Age         int    `bson:"age"`
    LastName    string `bson:"lastName"`
    FirstName   string `bson:"firstName"`
    Gender      string `bson:"gender"`
    VillageName string `bson:"villageName"`
}

//  delete the Applicant document in MongoDB based on a filter
func DeleteApplicant(filter bson.M) error {
    // MongoDB connection URI
    uri := "mongodb://localhost:27017"

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Set up MongoDB client
    clientOptions := options.Client().ApplyURI(uri)

    // Connect to MongoDB
    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        return err
    }
    defer client.Disconnect(ctx)

   
    database := client.Database("mydb") // Replace with your database name
    collection := database.Collection("aadhaar_card_applications")

    // Perform the delete operation
    _, err = collection.DeleteOne(ctx, filter)
    if err != nil {
        return err
    }

    return nil
}

func main() {
    // Define a filter to delete applicants with a certain age
    filter := bson.M{"age": 22}

    err := DeleteApplicant(filter)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Applicants with age 22 deleted successfully.")
}
