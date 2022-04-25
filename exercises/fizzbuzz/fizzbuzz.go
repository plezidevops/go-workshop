package main

import "fmt"

func main() {
	counter := 1

	for counter <= 20 {
		if counter%3 == 0 && counter%5 == 0 {
			fmt.Printf("count=%v,%v fizz buzz\n", counter, counter)
		} else if counter%3 == 0 {
			fmt.Printf("count=%v fizz\n", counter)
		} else if counter%5 == 0 {
			fmt.Printf("count=%v buzz\n", counter)
		} else {
			fmt.Println(counter)
		}
		counter++
	}
}
