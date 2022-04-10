package main

import "fmt"

// parameters and returns
func main() {
	var x int = 50
	var y int = 40
	var sum_value int

	// calling a function to get sum of values
	sum_value = sum(x, y)

	fmt.Printf("Sum of value is %d\n", sum_value)
}

// function returns the sum of two numbers
func sum(val1, val2 int) int {
	var result int
	result = val1 + val2
	return result
}
