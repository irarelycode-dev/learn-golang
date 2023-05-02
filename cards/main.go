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

	fmt.Println("***********************structs************************")

	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)

	var vignesh person
	fmt.Println(vignesh)
	fmt.Printf("%+v\n", vignesh)

	var myPerson person = person{
		firstName: "vignesh",
		lastName:  "pugazhendhi",
		contact:   contactInfo{email: "vickynsp15@gmail.com", zipCode: "H3H 2E7"},
	}
	myPerson.print()
	myPerson.updateName("vicky")
	myPerson.print()
	(&myPerson).updateFirstName("vicky")
	myPerson.print()

}
