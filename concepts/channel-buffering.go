package main

import "fmt"

func channelBuffering() {
	messages := make(chan string, 2)
	messages <- "chennai"
	messages <- "montreal"
	fmt.Print(<-messages)
	fmt.Print(<-messages)
}
