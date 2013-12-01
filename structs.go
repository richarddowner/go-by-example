/*
Typed collections of fields.
Useful for grouping data together
to form records
*/
package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {

	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})

	s := person{name: "Sean", age: 50} // Access struct fields with a dot
	fmt.Println(s.name)                // Sean

	sp := &s            // Can use dots with structs, pointers are automatically dereferenced
	fmt.Println(sp.age) // 50

	sp.age = 51         // stucts are mutable
	fmt.Println(sp.age) // 51
}
