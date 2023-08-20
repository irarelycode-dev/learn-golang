package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const numberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func pizzeria(pizzaMaker *Producer) {
	//keep a track of which pizza are we making
	var i = 0
	//run forever until we receive a quit notification
	//try to make pizzas
	for {
		currentPizza := makePizza(i)
		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
			case pizzaMaker.data <- *currentPizza:
			case quitChan := <-pizzaMaker.quit:
				close(pizzaMaker.data)
				close(quitChan)
				return
			}
		}
	}
}

func makePizza(i int) *PizzaOrder {
	i++
	if i <= numberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("Received order number #%d\n", i)
		rnd := rand.Intn(12) + 1
		var msg = ""
		var success = false
		if rnd > 8 {
			pizzasFailed++
		} else {
			pizzasMade++
		}
		total++
		fmt.Printf("Making pizza #%d. It will take %d seconds...\n", i, delay)

		//delay sometime
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("*** We ran out of ingredients for pizza #%d!", i)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** The cook quit while making pizza #%d!", i)
		} else {
			success = true
			msg = fmt.Sprintf("*** Pizza made for #%d!", i)
		}

		p := PizzaOrder{
			pizzaNumber: i,
			message:     msg,
			success:     success,
		}
		return &p

	}
	return &PizzaOrder{
		pizzaNumber: i,
	}
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func main() {

	//seed the random number generator
	rand.Seed(time.Now().UnixNano())
	color.Cyan("The pizzeria is open for business")
	color.Cyan("----------------------------------")
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	//run the producer in the background
	go pizzeria(pizzaJob)

}
