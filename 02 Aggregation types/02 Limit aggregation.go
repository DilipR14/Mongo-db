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

func performGroupAggregationWithLimit(limit int) ([]bson.M, error) {
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

    pipeline := mongo.Pipeline{
        bson.D{
            {"$group", bson.D{
                {"_id", "$villageName"},
                {"totalAge", bson.D{
                    {"$sum", "$age"},
                }},
                {"count", bson.D{
                    {"$sum", 1},
                }},
            }},
        },
        bson.D{
            {"$project", bson.D{
                {"_id", 0},
                {"Village", "$_id"},
                {"TotalAge", "$totalAge"},
                {"Count", "$count"},
                {"AverageAge", bson.D{
                    {"$divide", []interface{}{"$totalAge", "$count"}},
                }},
            }},
        },
        bson.D{
            {"$limit", limit}, // Add the $limit stage to limit the result
        },
    }

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
    limit := 5 // Set the limit to the number of documents you want
    aggregationResult, err := performGroupAggregationWithLimit(limit)
    if err != nil {
        log.Fatal(err)
    }

    for _, doc := range aggregationResult {
        fmt.Println(doc)
    }
}
