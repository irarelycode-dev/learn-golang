package main

import (
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

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func pizzeria(pizzaMaker *Producer) {

	//keep track of which pizza are making

	//run forever or until we receive a quit notification
	//try to make pizzas
	for {
		//try to make a pizza
		//decision
	}

}

func main() {

	rand.Seed(time.Now().UnixNano())
	color.Cyan("The pizzeria is open for business")
	color.Cyan("-----------------------------------")

	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	go pizzeria(pizzaJob)

}

// import (
// 	"fmt"
// 	"sync"
// )

// var wg sync.WaitGroup

// type Income struct {
// 	Source string
// 	Amount int
// }

// func main() {

// 	var bankBalance int
// 	var mutex sync.Mutex

// 	fmt.Printf("Intial bank account:$%d.00", bankBalance)
// 	fmt.Println()
// 	incomes := []Income{
// 		{Source: "Main job", Amount: 500},
// 		{Source: "Gifts", Amount: 10},
// 		{Source: "Part-time job", Amount: 50},
// 		{Source: "Investments", Amount: 100},
// 	}

// 	wg.Add(len(incomes))

// 	for i, income := range incomes {
// 		go func(i int, income Income) {
// 			defer wg.Done()
// 			for week := 1; week <= 52; week++ {
// 				mutex.Lock()
// 				temp := bankBalance
// 				temp += income.Amount
// 				bankBalance = temp
// 				mutex.Unlock()
// 				fmt.Printf("On week %d,you earned %d.00 from %s\n", week, income.Amount, income.Source)
// 			}
// 		}(i, income)
// 	}

// 	wg.Wait()
// 	fmt.Printf("Final bank balance:$%d.00\n", bankBalance)

// }

// import (
// 	"fmt"
// 	"sync"
// )

// var msg string
// var wg sync.WaitGroup

// func updateMessage(s string) {
// 	defer wg.Done()
// 	msg = s
// }

// func main() {
// 	msg = "hello world"
// 	wg.Add(2)

// 	go updateMessage("hello,universe!")
// 	go updateMessage("hello,cosmos!")

// 	wg.Wait()

// 	fmt.Println(msg)
// }

// func updateMessage(s string, mutex *sync.Mutex) {
// 	defer wg.Done()
// 	mutex.Lock()
// 	msg = s
// 	mutex.Unlock()
// }

// func main() {
// 	msg = "hello world"
// 	wg.Add(2)

// 	var mutex sync.Mutex

// 	go updateMessage("hello,universe!", &mutex)
// 	go updateMessage("hello,cosmos!", &mutex)

// 	wg.Wait()

// 	fmt.Println(msg)
// }
