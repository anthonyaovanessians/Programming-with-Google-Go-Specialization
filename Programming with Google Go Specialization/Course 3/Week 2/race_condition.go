/*

A race condition is a concurrency problem that may occur inside
a critical section. A critical section is a section of code that is
executed by multiple threads and where the sequence of execution
for the threads makes a difference in the result of the
concurrent execution of the critical section.

Whenever you run this program you can either get hello goodbye or goodbye hello printed
to the output.

The race condition in this program is due to
the interleaving and the interleaving is non-deterministic.
Therefore, it doesn't matter what the order of the goroutines are
because the outcome depends on non-deterministic ordering. Which means
the output could be hello goodbye or goodbye hello.

Since the outcome can either be hello goodbye or goodbye hello, we have a race condition.

*/

package main

import (
	"fmt"
	"time"
)

func hello() {

	fmt.Println("hello")
}

func goodbye() {

	fmt.Println("goodbye")
}

func main() {

	go hello()
	go goodbye()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("done")
}
