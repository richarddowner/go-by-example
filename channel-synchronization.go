/*
We can use channels to synchronize execution
across goroutines. This example uses a blocking
receive to wait for a goroutine to finish
*/
package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true // the done channel will notifiy another goroutine that this function's work is done
}

func main() {

	done := make(chan bool, 1)
	go worker(done) // start a worker goroutine giving it the channel to notify on

	<-done // block until we receive a notification from the worker on the channel
}
