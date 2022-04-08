package main

import "fmt"

// create a new type of deck
// which a slice of string

type deck []string // here we are defining a new type called deck that extends another type string

// the print function receives a deck (of cards) reference
func (d deck) print() { // receiver (d deck) on a function
	for i, card := range d {
		fmt.Println(i, card)
	}
}
