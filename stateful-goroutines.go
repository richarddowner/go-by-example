/*
This channel-based approach

In this example our state will be owned by a single
goroutine. This will guarantee that the data is never
corrupted with concurrent access. In order to read or write
that state, other goroutines will send messages to the owning
goroutine and receive corresponding replies. These readOp
and writeOp structs encapsulate those requests and a way
for the owning goroutine to respond.
*/
package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOp struct {
	key  int
	resp chan int
}
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {

	var ops int64 = 0

	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	// Here is the goroutine that owns the state, which is a map
	// as in the previous example but now private to the stateful
	// goroutine.
	//
	// This goroutine repeatedly selects on the reads and writes
	// channels, responding to requests as they arrive.
	//
	// A response is executed by first performing the requested
	// operation and then sending a value on the response channel
	// resp to indicate success.
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// Starts 100 goroutines to issue reads to the state-
	// owning goroutine via the reads channel. Each read
	// requires constucting a readOp, sending it over the
	// reads channel, and the receiving the result over the
	// provided resp channel
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := &readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddInt64(&ops, 1)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := &writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddInt64(&ops, 1)
			}
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)

}
