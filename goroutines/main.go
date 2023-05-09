package main

import (
	"fmt"
	"sync"
	"time"
)

type Person struct {
	name string
	age  int
}

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"gamma",
		"zeta",
		"peta",
		"theta",
		"pi",
		"epsilon",
	}

	wg.Add(len(words))

	for i, x := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, x), &wg)
	}

	wg.Wait()
	wg.Add(1)

	time.Sleep(1 * time.Second)
	printSomething("This is the 2nd thing to be printed", &wg)

	//channels,pointers concepts

	// tmp()

	// var channel chan int = make(chan int)

	// var customChannel chan Person = make(chan Person)

	// go tmp1(10, channel)
	// go tmp2(Person{name: "vignesh", age: 25}, customChannel)

	// var valueFromChannel int = <-channel
	// fmt.Println("valueFromChannel:", valueFromChannel)
	// fmt.Println("custom value from channel:", <-customChannel)

	// var sendOnlyChannel chan<- int = make(chan<- int)
	// fmt.Println(sendOnlyChannel)

	// var receiveOnlyChannel <-chan int = make(<-chan int)
	// fmt.Println(receiveOnlyChannel)

}
