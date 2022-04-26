package main

import "fmt"

func main() {
	a := 0

	for {
		if a > 9 {
			break
		}
		fmt.Printf("The value of a: %v\n", a)
		a++
	}
}
