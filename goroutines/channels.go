package main

func tmp1(value int, channel chan int) {
	channel <- value
}

func tmp2(value Person, channel chan Person) {
	channel <- value
}
