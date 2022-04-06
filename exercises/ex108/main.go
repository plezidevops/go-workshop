// changing multiple variable values in one line
package main

import "fmt"

func main() {

	// variable declaration
	query, limit, offset := "bat", 10, 0

	// changing the variables values
	query, limit, offset = "ball", offset, 20

	// print the variables content
	fmt.Println(query, limit, offset)

}
