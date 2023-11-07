package main

import (
	"errors"
	"fmt"
)

// Function that returns an error
func divide(a, b float64) (float64, error) {
	if b == 0 {
		// Package errors implements functions to manipulate errors.
		// The New function creates errors whose only content is a text message.
		err := errors.New("division by zero is not allowed")
		return 0, err
	}
	result := a / b
	return result, nil // No error, return result
}

func main() {
	// Attempt to divide two numbers
	numerator := 10.0
	denominator := 0.0

	result, err := divide(numerator, denominator)

	// Check if an error occurred
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Result: %.2f\n", result)
	}
}
