package main

import (
	"fmt"
)

// input: floating point number
// output: integer

var floatNum float64

func main() {

	truncValues()

}

func truncValues() {

	for i := 0; i < 2; i++ {

		fmt.Printf("Please enter a floating point number: ")

		_, err := fmt.Scan(&floatNum)
		if err != nil {
			panic("Please enter a floating point number.")
		}

		fmt.Printf("The truncated integer is: %v\n\n", int(floatNum))

	}

}
