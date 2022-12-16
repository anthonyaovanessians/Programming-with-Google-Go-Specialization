package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sli := make([]int, 0)
	for {
		if len(sli) < 10 {
			fmt.Println()
			fmt.Print("You will need to enter ten integers. Please enter one integer then press ENTER: ")
			reader := bufio.NewReader(os.Stdin)
			input, _ := reader.ReadString('\n')
			clean_input := strings.TrimSpace(input)
			//fmt.Println(sli)
			num, err := strconv.ParseInt(strings.TrimSpace(clean_input), 10, 64)
			if err != nil {
				fmt.Println("Value must be an integer.")
				continue
			} else {
				sli = append(sli, int(num))
				fmt.Println(sli)
			}

		} else {
			fmt.Println()
			fmt.Printf("Before Bubble Sorted slice: %v", sli)
			fmt.Println()
			BubbleSort(sli)
			break
		}

	}

}

func BubbleSort(sli []int) {
	for i := 0; i < len(sli); i++ {
		for j := 0; j < (len(sli)-i)-1; j++ {
			if sli[j] > sli[j+1] {
				Swap(sli, j)
			} else {
				//break
				continue
			}
		}

	}

	fmt.Println()
	fmt.Printf("After Bubble Sorted slice: %v"+"\n", sli)
	fmt.Println()
}

func Swap(sli []int, index int) {

	bigger := sli[index]
	lesser := sli[index+1]
	sli[index+1] = bigger
	sli[index] = lesser

}
