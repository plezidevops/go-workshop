// using var and shorthand to create an array

package main

import "fmt"

func creat_array() {
	var studentNames [10]string
	studentNames[0] = "Jude"
	fmt.Println(studentNames)
}

func shorthand() {
	studentNames := [5]string{"Hude", "Malpougra", "Mala", "Joel", "Joanito"}

	for i, name := range studentNames {
		fmt.Printf("%d - %s\n", i, name)
	}
}

func main() {
	creat_array()
	shorthand()
}
