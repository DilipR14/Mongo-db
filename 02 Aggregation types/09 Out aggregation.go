package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func performOutAggregation() error {
	uri := "mongodb://localhost:27017"

	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}
	defer client.Disconnect(context.Background())

	database := client.Database("mydb")
	collection := database.Collection("aadhaar_card_applications")

	// Define the aggregation pipeline with $out stage to write results to a new collection
	pipeline := mongo.Pipeline{
		bson.D{
			{"$match", bson.D{
				{"age", bson.D{
					{"$gte", 30}, // Match documents with age greater than or equal to 30
				}},
			}},
		},
		bson.D{
			{"$out", "filtered_aadhaar_data"}, // Create a new collection named "filtered_aadhaar_data"
		},
	}

	// Execute the aggregation with $out stage
	_, err = collection.Aggregate(context.Background(), pipeline)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := performOutAggregation()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Aggregation results written to the 'filtered_aadhaar_data' collection.")
}
