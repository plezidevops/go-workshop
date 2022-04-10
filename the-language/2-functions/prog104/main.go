// Call by value

package main

import "fmt"

func replace(x int) {
	x = 10
}

func main() {
	var x int = 10
	fmt.Printf("value of x before function call = %d", x)

	// call by value
	replace(x)
	fmt.Printf("\nvalue of x after function call = %d", x)

}
