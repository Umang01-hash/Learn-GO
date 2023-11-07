package main

import "fmt"

func main() {
	// Array
	fruits := [3]string{"apple", "banana", "cherry"}
	fruits[1] = "avocado"

	fmt.Println(fruits)

	// Slices are a more common and flexible way to work with sequences of data in Go as they don't have a fixed size
	// Slices are mutable, meaning you can add, remove, or modify elements.
	colors := []string{"red", "green", "blue"}

	fmt.Println(colors)

	// Use the append() function to add elements to a slice:
	colors = append(colors, "white")

	fmt.Println(colors)

	// Maps
	cityPopulations := map[string]int{
		"New York":    8419600,
		"Los Angeles": 3980000,
		"Chicago":     2716000,
	}

	fmt.Println(cityPopulations)

	// Structs are the data types that group together related fields of different types into a single unit
	type Person struct {
		FirstName string
		LastName  string
		Age       int
	}
	person1 := Person{"Alice", "Johnson", 28}

	fmt.Println(person1)
}
