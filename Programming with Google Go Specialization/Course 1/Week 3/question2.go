package main

import (
	"fmt"
)

func main() {
	x := [...]int{1, 2, 3, 4, 5}
	y := x[2:2]
	z := x[0:4]
	fmt.Print(len(x), cap(x), len(y), cap(y), len(z), cap(z))
}
