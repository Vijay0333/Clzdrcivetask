// package main

// import (
// 	"context"
// 	"fmt"
// 	"os"
// 	"time"

// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// func main() {
// 	// Load connection string from environment variable
// 	connectionString := os.Getenv("MONGO_CONNECTION_STRING")

// 	// Connect to MongoDB
// 	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connectionString))
// 	if err != nil {
// 		fmt.Println("Error connecting to MongoDB:", err)
// 		return
// 	}
// 	defer client.Disconnect(context.Background())

// 	// Get collection
// 	collection := client.Database("car_garage").Collection("cars")

// 	// Create car data
// 	carData := map[string]interface{}{
// 		"license_plate": "ABC123",
// 		"car_model":     "Honda Civic",
// 		"arrival_time":  time.Now(),
// 	}

// 	// Add car to database
// 	_, err = collection.InsertOne(context.Background(), carData)
// 	if err != nil {
// 		fmt.Println("Error adding car data:", err)
// 		return
// 	}

// 	fmt.Println("Car added successfully!")
// }
