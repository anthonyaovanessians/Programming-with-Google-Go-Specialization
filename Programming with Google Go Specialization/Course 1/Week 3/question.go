package main

import (
	"fmt"
)

func main() {
	x := [...]int{4, 8, 5}
	y := x[0:2]
	z := x[1:3]
	// y = {4, 8} 0,1
	// x = {8, 5} 0, 1
	y[0] = 1     //	change the zeroth element of y (which is four) to 1 {1, 8}
	z[1] = 3     // change the first element of z (which is five) to 3 {8, 3}
	fmt.Print(x) // final answer is {1, 8, 3} these commands change x because we are changing the underlying array?
}
