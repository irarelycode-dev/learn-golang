package main

import "fmt"

var deckSize int = 20

const a = 10

func main() {

	cards := newDeck()

	cards.print()
	fmt.Println("**************")

	hand, remainingCards := deal(cards, 5)

	hand.print()
	fmt.Println("**************")
	remainingCards.print()

	cards.saveToFile("my_cards")

	fmt.Println("******************reading from file and shuffling")

	cards = readDeckFromFile("my_cards")
	cards.print()
	fmt.Println("**************shuffle the cards")
	cards.shuffle()
	cards.print()

}
