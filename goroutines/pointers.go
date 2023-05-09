package main

import "fmt"

func tmp() {
	var p *int
	var i int = 45
	p = &i
	fmt.Println("Address of i:", p)
	fmt.Println("Value at address:", *p)
	*p = 21
	fmt.Println("Value at address:", *&*&i)
}
