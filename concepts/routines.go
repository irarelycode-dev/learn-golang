package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func channels() {
	messages := make(chan string)
	go func() { messages <- "ping" }()
	fmt.Println("reading from channel:", <-messages)
}

func goRoutines() {
	f("direct")
	go f("goroutine")
	go func(msg string) {
		fmt.Println(msg)
	}("going")
	time.Sleep(time.Second)
	channels()
}
