package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	sli := make([]int, 3)
	flag := true
	for flag {
		fmt.Println("Please enter an integer: ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input2 := strings.TrimSpace(input)
		switch input2 {
		case "X":
			flag = false
		default:
			num, err := strconv.ParseInt(strings.TrimSpace(input2), 10, 64)
			if err != nil {
				fmt.Println("Value must be an integer.")
				continue
			}
			sli = append(sli, int(num))
			sort.Ints(sli)
			fmt.Println(sli)

		}

	}

}
