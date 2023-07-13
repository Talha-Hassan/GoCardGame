package main

func main() {
	cards := newDeck()
	// // hand, remainingHand := cards.deal(5)
	// // hand.print()
	// // remainingHand.print()
	// // fmt.Println(cards.toString())
	// cards.saveToFile("Cards")
	// cards := newDeckFromFile("Cardss")c
	cards.shuffle()
	cards.print()
}

// array
// slice
