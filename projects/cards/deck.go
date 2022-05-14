package main

import "fmt"

// create a new type of deck
// which a slice of string

// here we are defining a new type called deck that sor t of extends another type string
type deck []string

// newDeck creating and returning a deck of cards
func newDeck() deck {
	// to hold the deck of cards
	cards := deck{}

	// list of cards and their values
	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// creating the card suites
	for _, suit := range cardSuites {
		for _, value := range cardValues {
			// add cards to the deck
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

// deal creates a hand of card
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// shuffle all the cards in the deck
func shuffle(d deck) deck {
	return d
}

// saves a deck to the hard drive
func saveFile
// loading a list od cards from the hard drive
