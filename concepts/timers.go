package main

import (
	"fmt"
	"time"
)

func timers() {
	fmt.Println("****************timers***********************")
	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C
	fmt.Println("This would run after:", 2)
	var timer2 = time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("timer 2 has been stopped")
	}
	fmt.Println("****************timers***********************")
}
