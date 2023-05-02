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

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p person) updateName(firstName string) {
	p.firstName = firstName
}

func (p *person) updateFirstName(firstName string) {
	(*p).firstName = firstName
}
