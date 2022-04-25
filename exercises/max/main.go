package main

import (
	"fmt"
)

func main() {
	num := []int{16, 8, 42, 4, 23, 15}
	max := 0

	for _, value := range num {
		if value > max {
			max = value
		}
	}
	fmt.Printf("The highest value is: %v", max)
}
