package main

import "fmt"

func main() {
	grades := [...]byte{'A', 'B', 'D', 'F', 'D', 'C'}

	// grade slice
	gPortion := grades[1:2]

	fmt.Println(gPortion)
	fmt.Println(len(gPortion))
	fmt.Println(cap(gPortion))

	slice := make([]string, 0, 10)
	fmt.Println(slice)
	fmt.Println(cap(slice))

}
