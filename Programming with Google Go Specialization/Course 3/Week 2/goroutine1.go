package main

import (
	"fmt"
	"time"
)

var balanceA int = 500
var balanceB int = 500

func Transfer1(balance1 *int, balance2 *int, amountToTransfer int) {
	if *balance1 > amountToTransfer {
		*balance1 -= amountToTransfer
		*balance2 += amountToTransfer
		fmt.Printf("Transfer 1: Balance A is %v \n", *balance1)
	} else {
		fmt.Println("Transfer 1: Error")
	}
}

func Transfer2(balance1 *int, balance2 *int, amountToTransfer int) {
	if *balance1 > amountToTransfer {
		*balance1 -= amountToTransfer
		*balance2 += amountToTransfer
		fmt.Printf("Transfer 2: Balance A is %v \n", *balance1)
	} else {
		fmt.Println("Transfer 2: Error")
	}
}

func main() {

	go Transfer1(&balanceA, &balanceB, 499)
	go Transfer2(&balanceA, &balanceB, 499)
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("done")
}
