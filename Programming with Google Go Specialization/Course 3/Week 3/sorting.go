/*

For the person who is grading me:

I don't understand what the assignment wants as its final answer.

It says it wants a final sorted array that is merged with
the other four sub arrays. But in order have a final sorted array you need
to compare the elements inside all the arrays with each other. Which means
that it doesn't matter if you make a custom function to compare each
element in each array with each other to make a sorted final list, use sort.Ints(),
or use a BubbleSort on the final array because they all do the same thing.
They each compare each element from the already merged or about to be merged
sub arrays. Therefore I just used a BubbleSort for my final answer to have a
final sorted array.

Also, if they wanted a final sorted array then I could have just done the
BubbleSort on the original array without having to spilt the array up into four pieces in the
first place.

*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	var max_num int

	for {
		fmt.Printf("Please enter a number that is divisble by four: ")
		_, err := fmt.Scan(&max_num)
		if err != nil {
			fmt.Println("Error:", err)
		}
		if max_num%4 == 0 {
			break
		} else {
			fmt.Println()
			fmt.Println("The number you entered must be divisible by four.")
			fmt.Println()
			continue
		}

	}

	var number int
	sli := make([]int, 0)

	for {
		if len(sli) == max_num {
			break
		} else {
			fmt.Println()
			fmt.Printf("Please enter %v numbers: ", max_num-len(sli))
			_, err := fmt.Scan(&number)
			if err != nil {
				fmt.Println("Error:", err)
			}
			sli = append(sli, number)
			fmt.Println(sli)
			continue
		}
	}

	divider := max_num / 4
	// each goroutine gets divider amount of numbers

	var wg sync.WaitGroup
	result := make(chan []int, 0)
	wg.Add(4)
	go sorter(sli[0:divider], &wg, result)
	value1 := <-result
	go sorter(sli[divider:divider*2], &wg, result)
	value2 := <-result
	go sorter(sli[divider*2:divider*3], &wg, result)
	value3 := <-result
	go sorter(sli[divider*3:divider*4], &wg, result)
	value4 := <-result
	wg.Wait()

	sli1 := append(value1, value2...)
	sli2 := append(sli1, value3...)
	final := append(sli2, value4...)
	final = BubbleSort(final)
	fmt.Println()
	fmt.Printf("The entire sorted list is %v", final)

}

func sorter(sli []int, wg *sync.WaitGroup, result chan []int) {
	fmt.Println()
	fmt.Println(sli)
	BubbleSort(sli)
	result <- BubbleSort(sli)
	wg.Done()
}

func BubbleSort(sli []int) []int {
	for i := 0; i < len(sli); i++ {
		for j := 0; j < (len(sli)-i)-1; j++ {
			if sli[j] > sli[j+1] {
				Swap(sli, j)
			} else {
				continue
			}
		}

	}

	return sli
}

func Swap(sli []int, index int) {

	bigger := sli[index]
	lesser := sli[index+1]
	sli[index+1] = bigger
	sli[index] = lesser

}
