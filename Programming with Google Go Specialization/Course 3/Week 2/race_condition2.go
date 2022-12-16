/*

Whenever you run this program you can either get a 1 or 2 printed
to the output.

The race condition in this program is due to
the interleaving and the interleaving is non-deterministic.
Therefore, it doesn't matter what the order of the goroutines are
because the outcome depends on non-deterministic ordering. Which means
the output could be a 1 or 2.

*/

package main

import (
	"fmt"
	"time"
)

//var x int = 1

func printer(num *int) {

	fmt.Println(*num)
}

func adder(num *int) {
	*num += 1
}

func main() {
	x := 1
	go printer(&x)
	go adder(&x)
	time.Sleep(3 * time.Second)
	fmt.Println("done")
}
