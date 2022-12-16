package main

import (
	"fmt"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (ani Animal) Eat() string {
	return ani.food
}

func (ani Animal) Move() string {
	return ani.locomotion
}

func (ani Animal) Speak() string {
	return ani.noise
}

var cow Animal = Animal{"grass", "walk", "moo"}
var bird Animal = Animal{"worms", "fly", "peep"}
var snake Animal = Animal{"mice", "slither", "hsss"}

func main() {

	var animal string
	var action string

	for {
		fmt.Print("> ")
		_, err := fmt.Scan(&animal, &action)
		if err != nil {
			fmt.Println("Error:", err)
		}

		animal = strings.ToLower(animal)
		action = strings.ToLower(action)

		switch animal {
		case "cow":
			switch action {
			case "eat":
				fmt.Println(cow.Eat())
			case "move":
				fmt.Println(cow.Move())
			case "speak":
				fmt.Println(cow.Speak())
			}

		case "bird":
			switch action {
			case "eat":
				fmt.Println(bird.Eat())
			case "move":
				fmt.Println(bird.Move())
			case "speak":
				fmt.Println(bird.Speak())
			}

		case "snake":
			switch action {
			case "eat":
				fmt.Println(snake.Eat())
			case "move":
				fmt.Println(snake.Move())
			case "speak":
				fmt.Println(snake.Speak())
			}

		default:
			fmt.Println("Nothing valid was picked.")

		}

	}

}
