package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func closure() {
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	tmp := intSeq()
	fmt.Println(tmp())
	fmt.Println(tmp())
}
