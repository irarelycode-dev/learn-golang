package main

import "fmt"

func channelConcepts() {
	myChannel := make(chan string)
	go func() {
		myChannel <- "ping something here in the channel"
	}()
	channelMsg := <-myChannel
	fmt.Println("Channel msg:", channelMsg)
}
