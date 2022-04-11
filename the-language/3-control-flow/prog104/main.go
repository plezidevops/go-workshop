// switch case

package main

import "fmt"

func main() {
	switch day := 2; day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Invalid")
	}
}
