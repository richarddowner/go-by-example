/*
Channels are the pipes that connect concurrent
goroutines. You can send values into channels
from one goroutine and receive those values into
another goroutine.

By default sends and receives block until both
the sender and receiver are ready. This property
allowed us to wait at the end of our program for
the 'ping' message without having to use any other
synchronization.
*/
package main

import (
	"fmt"
)

func main() {

	messages := make(chan string) // channels are typed by the values they convey

	go func() { messages <- "ping" }() // send a value into a channel using channel <-

	msg := <-messages // the <-channel receives a value from the channel
	fmt.Println(msg)
}
