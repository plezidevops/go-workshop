// using defer

package main

import "fmt"

func main() {
	defer fmt.Println("Go")
	fmt.Println("Book")
}
