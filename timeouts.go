/*
Important for programs that connect to
external resources or that otherwise need
to bound execution time.

Implementing timeouts in Go is easy and
elegant thanks to channels and select.
*/
package main

import (
	"fmt"
	"time"
)

func main() {

	// exec an external call that returns
	// its result on a channel c1 after 2s.
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	// res and <-time awaits the result
	// since res takes 2s, <-time returns first
	// timeout
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()

	// res and <-time awaits the result
	// since res takes 2s, res returns first
	// no timeout
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}
}
