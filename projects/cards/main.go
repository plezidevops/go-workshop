package main

func main() {
	//cards := newDeck()
	//cards.saveToFile("mycards.txt")
	//cards = newDeckFromFile("mycards.txt")
	//fmt.Println(cards)

	cards := newDeck()
	cards.shuffle()
	cards.print()
}
