package main

import (
	"fmt"
	"sync"
)

type Person struct {
	name string
	age  int
}

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func printSomethingDef(s string) {
	fmt.Println(s)
}

var msg string

func updateMessage(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	msg = s
}

func main() {

	// channels,pointers concepts
	//fmt.Println("******pointers and channels******")

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

	fmt.Println("********wait groups*************")

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

	wg.Wait() //this will wait till the wait group finishes
	wg.Add(1)

	// time.Sleep(1 * time.Second)
	go printSomething("This is the 2nd thing to be printed", &wg)

	fmt.Println("*******go routines assignment*****")
	msg = "Hello world!"
	wg.Add(1)
	go updateMessage("Hello universe!", &wg)
	wg.Wait()
	printSomethingDef(msg)

	wg.Add(1)
	go updateMessage("Hello,cosmos!", &wg)
	wg.Wait()
	printSomethingDef(msg)

	wg.Add(1)
	go updateMessage("Hello,world!", &wg)
	wg.Wait()
	printSomethingDef(msg)

	// fmt.Println("*******************go routines docs**************")
	// wg.Add(1)

	// go printSomething("You fucking vile being", &wg)

	// wg.Wait()
	// fmt.Println("Done working!")

	// fmt.Println("Anonymous function")

	// wg.Add(1)
	// go func() {
	// 	fmt.Println("Running anonymous function")
	// 	wg.Done()
	// }()
	// wg.Wait()
	// fmt.Println("Done executing!")

}
