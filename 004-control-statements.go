package main

import "fmt"

func main() {
	// If Statement
	age := 20
	if age >= 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a minor.")
	}

	// If with short statement
	if x := 42; x < 0 {
		fmt.Println("x is negative")
	}

	// For Loop
	for i := 1; i <= 5; i++ {
		fmt.Printf("Iteration %d\n", i)
	}

	// Switch Statement
	day := "Monday"
	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It's a weekday.")
	case "Saturday", "Sunday":
		fmt.Println("It's a weekend.")
	default:
		fmt.Println("Invalid day.")
	}
}
