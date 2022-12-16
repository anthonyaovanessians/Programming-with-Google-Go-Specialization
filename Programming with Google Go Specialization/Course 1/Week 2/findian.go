package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// input string
// output string

func main() {

	ianValues()

}

func ianValues() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Please enter a string: ")
	operator, _ := reader.ReadString('\n')

	clean_input := strings.TrimSpace(operator)
	lower_string := strings.ToLower(clean_input)
	end_index := len(lower_string)
	middle := lower_string[1 : end_index-1]

	if strings.HasPrefix(lower_string, "i") == true && strings.Contains(middle, "a") == true && strings.HasSuffix(lower_string, "n") == true {
		fmt.Printf("Found!")
	} else {
		fmt.Printf("Not Found!")
	}

}
