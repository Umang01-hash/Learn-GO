package main

import "fmt"

// Function declaration
func greet(name string) string {
	return "Hello, " + name + "!"
}

// Function with multiple parameters
func calculateArea(length float64, width float64) float64 {
	return length * width
}

// Function with return values
func findMax(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}

// Variadic function to calculate the sum of integers
func calculateSum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	// Function call
	message := greet("Alice")
	fmt.Println(message)

	area := calculateArea(10, 5)
	fmt.Println("Area:", area)

	maxNum := findMax(12, 8)
	fmt.Println("Maximum number:", maxNum)

	// Calling the variadic calculateSum function
	numbers := []int{1, 2, 3, 4, 5}
	total := calculateSum(numbers...)
	fmt.Printf("Total: %d\n", total)
}
