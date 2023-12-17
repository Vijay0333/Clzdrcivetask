package main

import (
	"context" // Import context package for context.Background
	"fmt"
	"io/ioutil"
	"net/http" // Import http package for status codes
	"os"   // Import os package for os.Getenv
	"time" // Import time package for time.Now
    "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gofr.dev/pkg/gofr"
)

func main() {
	// initialise gofr object
	app := gofr.New()

	// Serve index.html from the Template folder
	 app.GET("/", func(c *gofr.Context) (interface{}, error) {
		html, err := ioutil.ReadFile("Template/index.html")
		if err != nil {
			return nil, err
		}

		// Set proper response headers
		c.Response().Header().Set("Content-Type", "text/html")
		c.Response().WriteHeader(http.StatusOK)

		// Write HTML content to the response
		c.Response().Write(html)

		// Return nil and nil error to indicate successful response handling
		return nil, nil
	})

	// app.GET("/", func(c *gofr.Context) (interface{}, error) {
	// 	html, err := ioutil.ReadFile("Template/index.html")
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	return string(html), nil
	// })

	connectionString := os.Getenv("MONGO_CONNECTION_STRING")

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connectionString))
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return
	}
	defer client.Disconnect(context.Background())

	// Get collection
	collection := client.Database("car_garage").Collection("cars")

	// Create car data
	carData := map[string]interface{}{
		"license_plate": "ABC123",
		"car_model":     "Honda Civic",
		"arrival_time":  time.Now(),
	}

	// Add car to database
	_, err = collection.InsertOne(context.Background(), carData)
	if err != nil {
		fmt.Println("Error adding car data:", err)
		return
	}

	fmt.Println("Car added successfully!")

	// register route greet
	// app.GET("/greet", func(ctx *gofr.Context) (interface{}, error) {

	//     return "Hello World!", nil
	// })

	// Starts the server, it will listen on the default port 8000.
	// it can be over-ridden through configs
	app.Start()
}
