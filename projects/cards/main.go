package main

func main() {
	cards := newDeck()

	// cards.print() //call the print function and passed deck reference to it

	hand, remainingDeck := deal(cards, 5)

	hand.print()
	remainingDeck.print()

}
