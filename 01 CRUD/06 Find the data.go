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

func FindApplicants(filter bson.M) ([]Applicant, error) {
    // MongoDB connection URI
    uri := "mongodb://localhost:27017"

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Set up MongoDB client
    clientOptions := options.Client().ApplyURI(uri)

    // Connect to MongoDB
    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        return nil, err
    }
    defer client.Disconnect(ctx)

    
    database := client.Database("mydb") // Replace with your database name
    collection := database.Collection("aadhaar_card_applications")

    // Perform the find operation
    cursor, err := collection.Find(ctx, filter)
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    var applicants []Applicant
    for cursor.Next(ctx) {
        var applicant Applicant
        if err := cursor.Decode(&applicant); err != nil {
            return nil, err
        }
        applicants = append(applicants, applicant)
    }

    if err := cursor.Err(); err != nil {
        return nil, err
    }

    return applicants, nil
}

func main() {
    // Define a filter to find applicants with a certain age
    filter := bson.M{"age": 22}

    applicants, err := FindApplicants(filter)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Found applicants:")
    for _, applicant := range applicants {
        fmt.Printf("UAN: %s, Name: %s %s, Age: %d\n", applicant.UAN, applicant.FirstName, applicant.LastName, applicant.Age)
    }
}
