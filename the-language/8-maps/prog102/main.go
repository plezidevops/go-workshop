/*
	A slice has both a length and a capacity:
		The length of a slice is the number of elements it contains.
		The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
*/

package main

import "fmt"

func main() {
	// make creates a slice with 10 elements
	slice := make([]int, 0, 10)

	// the length of the slice is 0
	fmt.Println("Length of slice: ", len(slice))

	// The slice capacity is 10
	fmt.Println("The capacity of the slice: ", cap(slice))

	for i := 0; i < 10; i++ {
		slice = append(slice, i)
	}

	// the length of the slice is 10
	fmt.Println("Length of slice: ", len(slice))

	// The slice capacity is 10
	fmt.Println("The capacity of the slice: ", cap(slice))

	slice5 := slice[5:]
	fmt.Println("The length of slice5 is: ", len(slice5))
	fmt.Println("The capacity of slice5 is: ", cap(slice5))

	slice3 := slice5[:1]
	fmt.Println("The length of slice5 is: ", len(slice3))
	fmt.Println("The capacity of slice5 is: ", cap(slice3))
}
