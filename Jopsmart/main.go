package main

import (
	"context"
	"fmt"
	"gofr.dev/pkg/gofr"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"os"
	"time"
	"io/iouti"
)

// Car struct representing the data model
type Car struct {
	LicensePlate string    `json:"license_plate"`
	CarModel     string    `json:"car_model"`
	ArrivalTime  time.Time `json:"arrival_time"`
}

func main() {
	// Initialise gofr object
	app := gofr.New()

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

	// MongoDB connection string
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

	// Route to handle form submission
	app.POST("/submit", func(c *gofr.Context) (interface{}, error) {
		// Parse form data from the request
		err := c.Request.ParseForm()
		if err != nil {
			return nil, err
		}

		// Extract form values
		licensePlate := c.Request.FormValue("license_plate")
		carModel := c.Request.FormValue("car_model")
		arrivalTimeString := c.Request.FormValue("arrival_time")

		// Parse arrival time from form value
		arrivalTime, err := time.Parse("2006-01-02T15:04", arrivalTimeString)
		if err != nil {
			return nil, err
		}

		// Create Car object
		newCar := Car{
			LicensePlate: licensePlate,
			CarModel:     carModel,
			ArrivalTime:  arrivalTime,
		}

		// Insert the car data into MongoDB
		_, err = collection.InsertOne(context.Background(), newCar)
		if err != nil {
			return nil, err
		}

		// Return success message
		return "Car data stored in MongoDB", nil
	})

	// Start the server, it will listen on the default port 8000.
	// It can be overridden through configs
	app.Start()
}
