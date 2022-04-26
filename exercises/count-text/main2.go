package main

import (
	"fmt"
	"strings"
)

func main() {

	// text to work on
	text := "Go is a next-generation, open-source programming language created by Google, prized for its concurrency and connectivity. Using Go, developers can build modern applications that can actually save companies money on backend resources. This course is designed to help developers be productive with Go, starting with the essentials of the syntax. Learn the basics of Go basic types such as numbers and strings; working with conditionals and loops; creating object-oriented code with structs and methods; and handling errors. Instructor Miki Tebeka also emphasizes the concurrency features such as go routines and channels, and connectivity features for networking with APIs and databases. For the final project, Mika shows you how to build a highly concurrent server that combines everything you've learned into one elegant solution powered by Go."

	words := strings.Fields(text)

	// look for these punctuations
	punctuationMarks := []string{";", ".", ",", "'"}

	// map to hold occurrence of string
	counts := map[string]int{}

	// remove the punctuation marks
	for _, mark := range punctuationMarks {
		for j, word := range words {
			before, _, found := strings.Cut(word, mark)
			if found {
				words[j] = before
			}
		}
	}

	// loop through array and count words occurrence
	for _, value := range words {
		counts[strings.ToLower(value)]++

	}

	fmt.Println(counts)
}
