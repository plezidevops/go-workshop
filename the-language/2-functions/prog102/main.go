package main

import "fmt"

// multiple returns

// function that returns two values
func testFunc(x, y int) (int, int) {
	return x - y, x * y
}

func main() {
	var testvar1, testvar2 = testFunc(20, 30)

	fmt.Printf(("Addition result: %d\n"), testvar1)
	fmt.Printf("Multiplication result: %d\n", testvar2)
}
