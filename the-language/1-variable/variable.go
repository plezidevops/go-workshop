package main

import "fmt"

func main() {
	// Declare variables that are set to zero value
	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("var a int \t %T [%v]\n", a, a)
	fmt.Printf("var b int \t %T [%v]\n", b, b)
	fmt.Printf("var c int \t %T [%v]\n", c, c)
	fmt.Printf("var d int \t %T [%v]\n\n", d, d)

	// Declare variable and initialize it
	// We are using the short variable declaration operator

	aa := 10
	bb := "hello"
	cc := 89.90
	dd := true

	fmt.Printf("aa := 10 \t %T [%v]\n", aa, aa)
	fmt.Printf("	bb := \"hello\" \t %T [%v]\n", bb, bb)
	fmt.Printf("var c int \t %T [%v]\n", cc, cc)
	fmt.Printf("var d int \t %T [%v]\n\n", dd, dd)

}
