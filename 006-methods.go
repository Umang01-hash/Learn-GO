package main

import "fmt"

// Number is  a simple struct type
type Number struct {
	Value int
}

// Method to increment the value of number. This method has a receiver of type Number
func (n Number) incrementValue() {
	n.Value++
}

// Method with a pointer receiver. doubleValue can modify the value of the original struct Number
func (n *Number) doubleValue() {
	n.Value *= 2
}

func main() {
	// Create a Number instance
	num := Number{Value: 5}

	// Call the method with a value receiver to increment the value
	num.incrementValue()
	fmt.Printf("Value after increment: %d\n", num.Value)

	// Call the method with a pointer receiver to double the value
	num.doubleValue()
	fmt.Printf("Value after doubling: %d\n", num.Value)
}
