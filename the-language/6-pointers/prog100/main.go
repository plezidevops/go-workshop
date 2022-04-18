package main

import "fmt"

func main() {
	x := 1
	p := &x		// p points to the address of x
	fmt.Printf("The address of x is %v\n", p)
	fmt.Printf("The value of x is %v", *p)

	*p = 2
	fmt.Printf("The address of x is %v\n", p)
	fmt.Printf("The value of x is %v", *p)
}