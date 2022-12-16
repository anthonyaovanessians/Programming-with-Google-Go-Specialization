package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {

	var map1 = make(map[string]string)

	fmt.Println("Please enter your name: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input2 := strings.TrimSpace(input)
	map1["name"] = input2

	fmt.Println("Please enter your address: ")
	reader2 := bufio.NewReader(os.Stdin)
	input3, _ := reader2.ReadString('\n')
	input4 := strings.TrimSpace(input3)
	map1["address"] = input4

	//fmt.Println(map1)

	barr, err := json.Marshal(map1)
	if err != nil {
		fmt.Println("Error with json marshaling.")
	}

	fmt.Println(string(barr))

}
