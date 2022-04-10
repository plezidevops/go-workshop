// Variadic function

package main

import (
	"fmt"
	"strings"
)

// variadiv function to join strings
func join(element ...string) string {
	fmt.Println(element)

	return strings.Join(element, "_")
}

func main() {
	fmt.Println(join("GO", "Language", "Book"))
}
