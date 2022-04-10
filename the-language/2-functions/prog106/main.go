// variadic function

package main

import (
	"fmt"
	"strings"
)

func join(element ...string) string {
	return strings.Join(element, "_")
}

func main() {
	element := []string{"GO", "Programming", "Language"}
	fmt.Println(join(element...))
}
