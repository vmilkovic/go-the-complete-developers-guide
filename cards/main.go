package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.saveToFile("my_cards")
	newDeck := newDeckFromFile("my_cards")
	newDeck.print()
}
