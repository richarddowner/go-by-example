/*
Defer is used to ensure that a function call is
performed later in a program's execution.
Defer is often used where ensure and finally
would be used in other languages
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	f := createFile("C:/dev/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}
