package main

import (
	"gofr.dev/pkg/gofr"
	"github.com/Vijay0333/Jopsmart/models"
	"github.com/Vijay0333/Jopsmart/Database"
)

// ... (rest of your code remains unchanged)


func main() {
	// initialise gofr object
	app := gofr.New()

	app.POST("/cars", createCar)
	app.GET("/cars", getCars)
	app.GET("/cars/:id", getCarByID)
	app.PUT("/cars/:id", updateCar)
	app.DELETE("/cars/:id", deleteCar)

	// register route greet
	app.GET("/greet", func(ctx *gofr.Context) (interface{}, error) {

		return "./Database.go", nil
	})

	func createCar(ctx *gofr.Context) (interface{}, error) {
		// Parse car data from request body
		var car models.Car
		if err := ctx.BindJSON(&car); err != nil {
		  return nil, err
		}

		 // Store car in database
		 if err := database.CreateCar(&car); err != nil {
			return nil, err
		  }
		
		  return car, nil
		}
	  

	// Starts the server, it will listen on the default port 8000.
	// it can be over-ridden through configs
	app.Start()
}
