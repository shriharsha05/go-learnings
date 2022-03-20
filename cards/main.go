package main

func main() {

	cards := newDeck()

	// hand, remainingCards := deal(cards, 7)

	// hand.print()
	// remainingCards.print()

	// fmt.Println(hand.toString())

	// cards.saveToFile("my_cards")

	// cards := newDeckFromFile("my_cards")
	// cards.print()

	cards.shuffle()
	cards.print()
}
