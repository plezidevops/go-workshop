package main

import "fmt"

func variadic(x ...string) {
	fmt.Println(x[1])
	fmt.Println(x[3])
}

func main() {
	variadic("Go", "for", "the", "gold", "baby")
}
