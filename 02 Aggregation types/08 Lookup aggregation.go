package main

import (
    "context"
    "fmt"
    "log"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
)

func performLookupAggregation() ([]bson.M, error) {
    uri := "mongodb://localhost:27017"

    clientOptions := options.Client().ApplyURI(uri)

    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        return nil, err
    }
    defer client.Disconnect(context.Background())

    err = client.Ping(context.Background(), readpref.Primary())
    if err != nil {
        return nil, err
    }

    database := client.Database("mydb")
    collection := database.Collection("aadhaar_card_applications")

    // Define the main aggregation pipeline
    pipeline := mongo.Pipeline{
        bson.D{
            {"$lookup", bson.D{
                {"from", "another_collection"}, // Replace with the name of the other collection
                {"localField", "UAN"},          // Field from the input collection (aadhaar_card_applications)
                {"foreignField", "UAN"},       // Field from the other collection
                {"as", "joinedData"},           // Alias for the joined data
            }},
        },
    }

    // Execute the lookup aggregation
    cursor, err := collection.Aggregate(context.Background(), pipeline)
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.Background())

    var result []bson.M
    if err := cursor.All(context.Background(), &result); err != nil {
        return nil, err
    }

    return result, nil
}

func main() {
    lookupResult, err := performLookupAggregation()
    if err != nil {
        log.Fatal(err)
    }

    for _, doc := range lookupResult {
        fmt.Println(doc)
    }
}
