package main

func main() {
	cards := newDeck()
	cards = cards.shuffleCards()
	cards.printDeck()
}
