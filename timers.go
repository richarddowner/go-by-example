/*
Useful for executing Go code at
some point in the future.
*/
package main

import (
	"fmt"
	"time"
)

func main() {

	timer1 := time.NewTimer(time.Second * 2)

	// blocks on the timer's channel C until it sends
	// a value indicating that the timer expired
	<-timer1.C
	fmt.Println("Timer 1 expired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 { // Will stop before expires
		fmt.Println("Timer 2 stopped")
	}
}
