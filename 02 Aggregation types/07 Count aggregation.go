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

func performCountAggregation(minAge int) (int64, error) {
    uri := "mongodb://localhost:27017"

    clientOptions := options.Client().ApplyURI(uri)

    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        return 0, err
    }
    defer client.Disconnect(context.Background())

    err = client.Ping(context.Background(), readpref.Primary())
    if err != nil {
        return 0, err
    }

    database := client.Database("mydb")
    collection := database.Collection("aadhaar_card_applications")

    pipeline := mongo.Pipeline{
        bson.D{
            {"$match", bson.D{
                {"age", bson.D{
                    {"$gte", minAge},
                }},
            }},
        },
        bson.D{
            {"$count", "totalDocuments"}, // Count the number of documents that match the criteria
        },
    }

    cursor, err := collection.Aggregate(context.Background(), pipeline)
    if err != nil {
        return 0, err
    }
    defer cursor.Close(context.Background())

    var result struct {
        TotalDocuments int64 `bson:"totalDocuments"`
    }

    if cursor.Next(context.Background()) {
        err := cursor.Decode(&result)
        if err != nil {
            return 0, err
        }
    }

    return result.TotalDocuments, nil
}

func main() {
    minAge := 30 // Minimum age for filtering

    count, err := performCountAggregation(minAge)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Total number of documents with age greater than or equal to %d: %d\n", minAge, count)
}
