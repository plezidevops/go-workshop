package main

import "fmt"

// create a new type of deck
// which a slice of string

type deck []string // here we are defining a new type called deck that extends another type string

func newDeck() deck { // Creating a deck of cards
	cards := deck{}

	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// the print function receives a deck (of cards) reference
func (d deck) print() { // receiver (d deck) on a function
	for i, card := range d {
		fmt.Println(i, card)
	}
}
