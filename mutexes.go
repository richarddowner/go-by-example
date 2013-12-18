/*
Safely access data across multiple goroutines.
*/
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	var state = make(map[int]int)

	var mutex = &sync.Mutex{}

	var ops int64 = 0

	// for each read we pick a key to access,
	// Lock() the mutex to ensure exclusive access
	// to the state, read the value at the chosen key,
	// Unlock() the mutex, and increment the ops count
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {

				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)

				// In order to ensure that this goroutine
				// doesn't starve the scheduler, we explicity
				// yeild after each operation with runtime.Gosched()
				runtime.Gosched()
			}
		}()
	}

	// Start 1- goroutines to simulate writes,
	// using the same pattern we did for reads
	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)

	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}
