package main

import "fmt"

func main() {
	age := 23
	fmt.Println(age)

	myAge := &age
	fmt.Println(*myAge)

	*myAge = 44
	fmt.Println(*myAge)
	fmt.Println(age)

	n1 := 4
	n2 := 5

	var num1 *int = &n1
	var num2 *int = &n2

	fmt.Printf("num1 = %d and num2 = %d\n", *num1, *num2)
	result := add(num1, num2)

	fmt.Printf("num1 = %d and num2 = %d\n", *num1, *num2)
	fmt.Printf("result = %d", result)

}

func add(num1 *int, num2 *int) int {
	*num1 = 6
	*num2 = 3
	return *num1 + *num2
}
