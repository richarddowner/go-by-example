/*
The primary mechanism for managing state
in Go is communication over channels.
There are a few other options for managing
state though.
Here we look at using the sync/atomic package
for atomic counters accessed by multiple
goroutines.
*/
package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func main() {

	var ops uint64 = 0

	for i := 0; i < 50; i++ {
		go func() {
			for {

				atomic.AddUint64(&ops, 1)

				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
