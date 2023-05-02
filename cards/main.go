package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

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

	var a map[string]string = map[string]string{
		"red": "0",
	}
	fmt.Println(a)
	for key, value := range a {
		fmt.Println(key, value)
	}
	delete(a, "red")
	fmt.Println(a)

	fmt.Println("***********************interfaces************************")

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

	fmt.Println("********Http package**************")

	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println(resp)

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	io.Copy(os.Stdout, resp.Body)

}
