package main

import "fmt"

// show the use of name return arguments

func sum(a, b int) (add int, sub int) {
	add = a + b
	sub = b - a

	return
}

func main() {
	x, y := sum(10, 20)

	fmt.Printf("10 + 20 = %d, %d\n", x, y)
}
