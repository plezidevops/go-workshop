// program to print a radom number between 1 and 5 to the console
package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {

	// set random seed generator
	rand.Seed(time.Now().UnixNano())

	// generate a random number
	r := rand.Intn(5) + 1

	// repat * r times
	stars := strings.Repeat("*", r)

	// print the string
	fmt.Println(stars)
}
