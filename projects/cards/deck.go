package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

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

func (d deck) toString() string {
	return strings.Join(d, ",")
}

// saves a deck to the hard drive
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0644)
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

// shuffle all the cards in the deck
func (d deck) shuffle() {
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}
