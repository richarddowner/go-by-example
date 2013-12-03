/*
A goroutine is a lightweight thread of execution
*/
package main

import (
	"fmt"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	f("direct")

	go f("goroutine") // will execute concurrently with calling one

	go func(msg string) { // go routine on anonymous function call
		fmt.Println(msg)
	}("going")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
