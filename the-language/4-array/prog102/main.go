// array length 3D array

package main

import "fmt"

func arrayLength() {
	a := [...]int{90, 34, 67, 23}

	fmt.Println(len(a))
}

// showing the position of element in an array
func arrayElementPosition() {
	a := [...]int{0: 34, 9: 23, 2: 12, 5: 78}
	fmt.Println(a)
	fmt.Printf("Element at position 9 is %d\n", a[1])
}

// 3D array
func studentInfo() {
	names := [3][3]string{
		{"jude", "math", "B"},
		{"pierre", "statistics", "A"},
		{"paul", "calculus", "B+"},
	}

	fmt.Println(names[0][0])
	fmt.Println(names[0][1])
	fmt.Println(names[0][2])
}

func main() {
	arrayLength()
	arrayElementPosition()
	studentInfo()
}
