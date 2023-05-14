package main

import (
	"fmt"
	"net/http"
)

var deckSize int = 20

const a = 10

func main() {

	fmt.Println("********************channels and go routines**********************")
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	channel := make(chan string)

	for _, link := range links {
		go checkLink(link, channel)
	}
	for {
		go checkLink(<-channel, channel)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "link might be down")
		c <- link
		return
	}

	fmt.Println(link, " is up!")
	c <- link

}
