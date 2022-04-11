// if else statement

package main

import (
	"fmt"
	"strings"
)

func main() {
	var x string = "Port-Aux-Prince"

	if strings.ToLower(x) == "port-aux-prince" {
		fmt.Printf(" %v is the of Capital of Haiti", x)
	} else {
		fmt.Printf("%v is not the capital of Haiti.", x)
	}
}
