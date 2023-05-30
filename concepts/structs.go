package main

import "fmt"

type Person struct {
	name string
	age  int
}

func newPerson(name string) *Person {
	p := Person{name: name}
	p.age = 42
	return &p
}

func structFunc() {
	tmp := newPerson("vignesh")
	fmt.Println("name:", tmp.age, tmp.name)
	fmt.Println("person:", tmp)
}
