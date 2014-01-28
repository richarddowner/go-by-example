package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

var pf = fmt.Printf

func main() {

	p := point{1, 2}

	pf("%v\n", p) // {1 2}

	pf("%+v\n", p) // {x:1, y:2}

	pf("%#v\n", p) // main.point{x:1, y:2}

	pf("%T\n", p) // main.point

	pf("%t\n", true) // true

	pf("%d\n", 123) // 123

	pf("%b\n", 14) // 1110

	pf("%c\n", 33) // !

	pf("%x\n", 456) // 1c8

	pf("%f\n", 78.9) // 78.900000

	pf("%e\n", 123400000.0) // 1.234000e+08

	pf("%E\n", 123400000.0) // 1.234000E+08

	pf("%s\n", "\"string\"") // "string"

	pf("%q\n", "\"string\"") // "\"string\""

	pf("%x\n", "hex this") // 6865782074686973

	pf("%p\n", &p) // 0x42135100

	pf("|%6d|%6d|\n", 12, 345) // |    12|   345|

	pf("|%6.2f|%6.2f|\n", 1.2, 3.45) // |  1.20|  3.45|

	pf("|%-6.2f|%-6.2f|\n", 1.2, 3.45) // |1.20  |3.45  |

	pf("|%6s|%6s|\n", "foo", "b") // |   foo|     b|

	pf("|%-6s|%-6s|\n", "foo", "b") // |foo   |b     |

	s := fmt.Sprintf("a %s", "string")

	fmt.Println(s) // a string

	fmt.Fprintf(os.Stderr, "an %s\n", "error") // an error
}
