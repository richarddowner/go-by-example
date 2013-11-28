/*
Maps are Go's built-in associative data type
(sometimes called hashes or dicts in other langs)

To create an empty map, use the builtin make:
make(map[key-type]val-type).
*/
package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m) // map: map[k1:7 k2:13]

	v1 := m["k1"]
	fmt.Println("v1: ", v1)     // v1: 7
	fmt.Println("len:", len(m)) // len: 2

	delete(m, "k2")
	fmt.Println("map:", m) // map: map[k1:7]

	_, prs := m["k2"]
	fmt.Println("prs:", prs) // prs: false

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n) // map: map[foo:1 bar:2]
}
