package main

import "fmt"

func pointers() {
	var a int = 10
	var addr *int = &a
	fmt.Println("var a and it's address:", a, addr)
}
