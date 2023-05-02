package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode string
}

type englishBot struct{}

type spanishBot struct{}

type bot interface {
	getGreeting() string
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p person) updateName(firstName string) {
	p.firstName = firstName
}

func (p *person) updateFirstName(firstName string) {
	(*p).firstName = firstName
}

func (eb englishBot) getGreeting() string {
	return "hi,there!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
