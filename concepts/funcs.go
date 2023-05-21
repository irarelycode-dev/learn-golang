package main

import "fmt"

func vals() (int, int) {
	return 5, 7
}

func funcs() {
	a, b := vals()
	fmt.Println("vals func:", a, b)
}
