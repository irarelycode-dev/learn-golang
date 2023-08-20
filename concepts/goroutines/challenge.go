package main

import (
	"fmt"
	"sync"
)

var msg string

var wg sync.WaitGroup

func updateMessage(s string) {
	defer wg.Done()
	msg = s
}

func printMessage() {
	fmt.Println(msg)
}

func challenge() {
	msg = "Hello,world!"

	wg.Add(1)
	go updateMessage("Hello,universe!")
	wg.Wait()
	printMessage()

	wg.Add(1)
	go updateMessage("Hello,cosmos!")
	wg.Wait()
	printMessage()

	wg.Add(1)
	go updateMessage("Hello,galaxy!")
	wg.Wait()
	printMessage()
}
