package main

func main() {
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	cards.print() //call the print function and passed deck reference to it

}

func newCard() string {
	return "Five of Diamonds"
}
