// default package when is run when execute the program. By declaring it you're telling the Go compiler that this file is part of the main package
//
//	and it should be compiled into an executable binary, not a library.
//
// Compiling - way to turn code into smth that a computer can understand and execute. Two types: executable binaries and library.
// Executable binary is a proggram that you can run directly on your computer. This file contains all the instractions that your computer
// needs to execute tu run your program.
// In go everything is a package and we have to include "main package" in our applications
package main

import (
	"fmt"
	"math"
)

// Go is organized into packages. Go's standard library provides different core packages for us to use. We can use them by imprting them.
// we need to declare from which packages specific func are coming from.

// create an object with data types for each field in this object . user and restuarnats
type User struct {
	name    string
	latLong [2]float64
	// data type of methods [pickup: true, delivery: false] is a map with key string and value boolean
	methods string
}

type Restaurant struct {
	name        string
	description string
	category    string
	latLong     [2]float64
	// menu        []menu   // Use []menu for menu field
	menu    interface{}
	open    bool
	methods map[string]bool
}

func main() {
	// var users = []user{}
	// var restaurants = []restaurant{}

	// create a sample user
	user := User{
		// ID:      1,
		name:    "John Doe",
		latLong: [2]float64{37.7749, -122.4194},
		methods: "pickup",
	}

	restaurants := []Restaurant{
		{
			// id:          1,
			name:        "Burger King",
			description: "Fast food chain",
			category:    "burgers",
			latLong:     [2]float64{37.7749, -122.4194},
			menu:        []string{"Whopper", "Fries", "Soda"},
			open:        true,
			methods:     map[string]bool{"pickup": true, "delivery": false},
		},
		{
			// id:          2,
			name:        "McDonald's",
			description: "Fast food chain",
			category:    "burgers",
			latLong:     [2]float64{37.7819, -122.4324},
			menu:        []string{"Big Mac", "Fries", "Soda"},
			open:        false,
			methods:     map[string]bool{"pickup": false, "delivery": true},
		},
	}
	// add more restaurants here

	// get the matching restaurants for the user
	matchingRestaurants := getRestaurants(user, restaurants)
	fmt.Println(user.name, matchingRestaurants)
}

// create the function which will return matching restaurants, specify data type of each argument
func getRestaurants(user User, restaurants []Restaurant) []Restaurant {
	// create an empty slice where we can store matching restuarants
	matchingRestaurants := []Restaurant{}
	// uint is for negative and positive but use float as the distance will be in flat so the data type has to match
	var deliveryRadius float64 = 5.0

	// loop through all restaurants
	for _, restaurant := range restaurants {
		//check if the restaurant is open
		if restaurant.open {
			//if its open, check if the delivery method matches users pickup method
			if restaurant.methods[user.methods] == true {
				// calculate the distance between user and restaurant
				distance := calculateDistance(user.latLong, restaurant.latLong)
				// check if distance is equal to the delivery radius
				if distance == deliveryRadius {
					//add restaurants to matchingRestaurant slice
					matchingRestaurants = append(matchingRestaurants, restaurant)
				}
			}
		}
	}
	return matchingRestaurants
}

// always specify data type
// calculateDistance returns the distance in kilometers between two lat-long coordinates
func calculateDistance(location1, location2 [2]float64) float64 {
	lat1, lon1 := location1[0], location1[1]
	lat2, lon2 := location2[0], location2[1]
	radius := 6371 // Earth's radius in kilometers

	// convert latitude and longitude to radians
	lat1Rad := degreesToRadians(lat1)
	lon1Rad := degreesToRadians(lon1)
	lat2Rad := degreesToRadians(lat2)
	lon2Rad := degreesToRadians(lon2)

	// calculate the difference between the latitudes and longitudes
	dLat := lat2Rad - lat1Rad
	dLon := lon2Rad - lon1Rad

	// calculate the distance using the haversine formula
	a := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	distance := float64(radius) * c

	return distance
}

// degreesToRadians converts degrees to radians
func degreesToRadians(degrees float64) float64 {
	return degrees * (math.Pi / 180)
}

// we need to tell Go when to start execute the program. main func is an entry point of the application.

// func main(restaurants) restaurants {
// 	openRestaurants := []string
// 	for _, restuarant := range restaurants {
// 		// how to access attribute inside the object
// 		if open == true {
// 			openRestaurants = append(openRestaurants, restaurant)
// 		}
// 	}
// 	return openRestaurants
// }

// var restaurants = []string{"burher", }

// define the Restaurant struct
// type restaurant struct {
// 	name        string
// 	description string
// 	category    string
// 	latLong     []float64
// 	menu        interface{}
// 	open        bool
// 	// methods     map[string]bool
// }

// What to do if you don’t know the type, do you use an empty interface?

// How to access an attribute inside an object?

// Keep main clean, call another function from inside it.
// https://codesource.io/how-to-use-array-of-structs-in-golang/
// https://gobyexample.com/functions
