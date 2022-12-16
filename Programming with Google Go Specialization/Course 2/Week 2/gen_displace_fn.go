package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	sli := make([]float64, 0)
	for {
		fmt.Println()
		fmt.Println("Please enter acceleration value: ")
		reader := bufio.NewReader(os.Stdin)
		input_acceleration, _ := reader.ReadString('\n')
		clean_acceleration := strings.TrimSpace(input_acceleration)
		acceleration, err := strconv.ParseFloat(clean_acceleration, 64)
		if err != nil {
			fmt.Println("Value must be a float.")
			continue
		} else {
			sli = append(sli, float64(acceleration))
			break
		}

	}
	for {
		fmt.Println()
		fmt.Println("Please enter initial velocity value: ")
		reader := bufio.NewReader(os.Stdin)
		input_initial_velocity, _ := reader.ReadString('\n')
		clean_initial_velocity := strings.TrimSpace(input_initial_velocity)
		initial_velocity, err := strconv.ParseFloat(clean_initial_velocity, 64)
		if err != nil {
			fmt.Println("Value must be a float.")
			continue
		} else {
			sli = append(sli, float64(initial_velocity))
			break
		}

	}

	for {
		fmt.Println()
		fmt.Println("Please enter initial displacement value: ")
		reader := bufio.NewReader(os.Stdin)
		input_initial_displacement, _ := reader.ReadString('\n')
		clean_initial_displacement := strings.TrimSpace(input_initial_displacement)
		initial_displacement, err := strconv.ParseFloat(clean_initial_displacement, 64)
		if err != nil {
			fmt.Println("Value must be a float.")
			continue
		} else {
			sli = append(sli, float64(initial_displacement))
			break
		}

	}

	fmt.Println()
	acceleration := sli[0]
	initial_velocity := sli[1]
	initial_displacement := sli[2]

	fn := GenDisplaceFn(acceleration, initial_velocity, initial_displacement)

	for {
		fmt.Println("Please enter a time value: ")
		reader := bufio.NewReader(os.Stdin)
		input_time, _ := reader.ReadString('\n')
		clean_time := strings.TrimSpace(input_time)
		time, err := strconv.ParseFloat(clean_time, 64)
		if err != nil {
			fmt.Println("Value must be a float.")
			continue
		} else {
			fmt.Println()
			fmt.Printf("Displacement is %v", fn(time))
			break
		}

	}
}

func GenDisplaceFn(a, v_o, s_o float64) func(float64) float64 {
	displacement := func(t float64) float64 {
		return (float64(1)/float64(2))*a*math.Pow(t, 2) + v_o*t + s_o
	}

	return displacement
}
