package main

func main() {

	// cards := newDeck()
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")

	// cards.print()

	// hand, remainingCards := deal(cards, 5)

	// fmt.Println("Hand")
	// hand.print()
	// fmt.Println("Remaining Cards")
	// remainingCards.print()

	// cards := newDeckFromFile("my_cards")
	// cards.print()
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
