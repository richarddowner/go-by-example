/*
In previous examples we saw how for and range
provide iteration over basic data structures.
We can also use this syntax to iterate over
values recieved from a channel.
*/
package main

import (
	"fmt"
)

func main() {

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// iterate over channel
	for elem := range queue {
		fmt.Println(elem)
	}
}
