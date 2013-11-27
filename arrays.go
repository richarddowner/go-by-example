package main

import "fmt"

func main() {

	var a [5]int
	fmt.Println("emp:", a) // emp: [0 0 0 0 0]

	a[4] = 100
	fmt.Println("set:", a) // set: [0 0 0 0 100]
	fmt.Println("get:", a[4]) // get: 100

	fmt.Println("len:", len(a)) // len: 5

	a[0] = 0
	a[1] = 1
	a[2] = 2
	a[3] = 3
	a[4] = 4

	fmt.Println("sub:", a[0:5]) // sub: [0 1 2 3 4]
	fmt.Println("sub:", a[0:]) // sub: [0 1 2 3 4]
	fmt.Println("sub:", a[:5]) // sub: [0 1 2 3 4]
	fmt.Println("sub:", a[2:5]) // sub: [2 3 4] ~ inclusive:non-inclusive

	b := [5]int{1,2,3,4,5}
	fmt.Println("dcl:", b) // dcl: [1 2 3 4 5]

	var twoD [2][3]int
	for i:= 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD) // 2d: [[0 1 2] [1 2 3]]
}
