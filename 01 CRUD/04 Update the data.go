package main

import (
	"context"
	"fmt"
	"log"
	"time"

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

// inserts a single Applicant document into MongoDB
func InsertApplicant(applicant Applicant) error {
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

	// Insert the Applicant document
	_, err = collection.InsertOne(ctx, applicant)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// Create an Applicant document
	applicant := Applicant{
		UAN:         "7fe1e4ae-7be7-4557-a52b-65005ca65276",
		Age:         22,
		LastName:    "Ravi",
		FirstName:   "Kira",
		Gender:      "male",
		VillageName: "SRK",
	}

	// Call the InsertApplicant function to insert the document
	err := InsertApplicant(applicant)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Applicant inserted successfully.")
}
