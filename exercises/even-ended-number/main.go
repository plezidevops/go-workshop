/*
	An even-ended-number is a number with the same first and last digits. Write a program that count how many even-ended-numbers result from multiplying two four-digit numbers.
*/
package main

import "fmt"

func main() {

	count := 0

	firstNumber := 1000
	secondNumber := 9999

	for x := firstNumber; x <= secondNumber; x++ {
		for y := x; y <= secondNumber; y++ {
			n := x * y

			s := fmt.Sprintf("%d", n)

			if s[0] == s[len(s)-1] {
				count++
			}
		}
	}

	fmt.Printf("Number of times: %v", count)
}
