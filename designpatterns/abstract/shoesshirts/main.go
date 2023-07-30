package main

import "fmt"

func main() {
	adidasFactory, _ := GetSportsFactory("adidas")
	nikeFactory, _ := GetSportsFactory("nike")

	nikeShoe := nikeFactory.makeShoe()
	nikeShirt := nikeFactory.makeShirt()

	adidasShoe := adidasFactory.makeShoe()
	adidasShirt := adidasFactory.makeShirt()

	shoeDetails(nikeShoe)
	shirtDetails(nikeShirt)

	shoeDetails(adidasShoe)
	shirtDetails(adidasShirt)

}

func shoeDetails(s IShoe) {
	fmt.Println("***Shoe details***")
	fmt.Println("Logo:", s.getLogo())
	fmt.Println("Size:", s.getSize())
	fmt.Println("***Shoe details***")
}

func shirtDetails(s IShirt) {
	fmt.Println("***Shirt details***")
	fmt.Println("Logo:", s.getLogo())
	fmt.Println("Size:", s.getSize())
	fmt.Println("***Shirt details***")
}
