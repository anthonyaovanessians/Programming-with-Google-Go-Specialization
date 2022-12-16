package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	type Person struct {
		fname string
		lname string
	}

	fmt.Println("Please enter the name of your file: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input2 := strings.TrimSpace(input)

	dat, err := ioutil.ReadFile(input2)
	if err != nil {
		fmt.Println("File not found")
	}

	fmt.Println()

	arr := strings.SplitAfter(string(dat), "\n")

	sli := make([]Person, 0)
	n := 0

	for i := range arr {

		if i == len(arr) {
			break
		}

		arr_for := strings.Split(arr[i], " ")

		per := Person{
			fname: arr_for[n],
			lname: arr_for[n+1],
		}

		per.fname = addString(per.fname, 20)
		per.lname = addString(per.lname, 20)

		sli = append(sli, per)

	}

	for i := range sli {
		fmt.Println(sli[i].fname + sli[i].lname)
	}

	fmt.Println()
}

func addString(name string, length int) string {

	flag := true
	for flag {
		if len(name) != length {
			name += " "
			continue
		} else {
			break
		}
	}

	return name
}
