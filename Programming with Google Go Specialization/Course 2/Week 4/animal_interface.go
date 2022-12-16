package main

import (
	"fmt"
	"strings"
)

type Animal interface {
	Eat() string
	Move() string
	Speak() string
}

type cow struct {
	name       string
	food       string
	locomotion string
	sound      string
}

func (c cow) Eat() string {
	return c.food
}

func (c cow) Move() string {
	return c.locomotion
}

func (c cow) Speak() string {
	return c.sound
}

type bird struct {
	name       string
	food       string
	locomotion string
	sound      string
}

func (b bird) Eat() string {
	return b.food
}

func (b bird) Move() string {
	return b.locomotion
}

func (b bird) Speak() string {
	return b.sound
}

type snake struct {
	name       string
	food       string
	locomotion string
	sound      string
}

func (s snake) Eat() string {
	return s.food
}

func (s snake) Move() string {
	return s.locomotion
}

func (s snake) Speak() string {
	return s.sound
}

var made_cow = cow{"", "grass", "walk", "moo"}
var made_bird = bird{"", "worms", "fly", "peep"}
var made_snake = snake{"", "mice", "slither", "hsss"}

func main() {

	for {
		var command1 string
		var command2 string
		var command3 string

		fmt.Print("> ")
		fmt.Scan(&command1, &command2, &command3)

		command1 = strings.ToLower(command1)
		command2 = strings.ToLower(command2)
		command3 = strings.ToLower(command3)

		if command1 == "newanimal" {
			switch command3 {
			case "cow":
				made_cow = cow{command2, "grass", "walk", "moo"}
				fmt.Println("Created it!")
			case "bird":
				made_bird = bird{command2, "worms", "fly", "peep"}
				fmt.Println("Created it!")
			case "snake":
				made_snake = snake{command2, "mice", "slither", "hsss"}
				fmt.Println("Created it!")
			default:
				fmt.Printf("%v is not a valid animal.\n", command3)
			}

		} else if command1 == "query" {
			if command2 != " " {
				switch command2 {
				case made_cow.name:
					switch command3 {
					case "eat":
						fmt.Println(made_cow.Eat())
					case "move":
						fmt.Println(made_cow.Move())
					case "speak":
						fmt.Println(made_cow.Speak())
					default:
						fmt.Printf("%v is not a valid command.\n", command3)
					}
				case made_bird.name:
					switch command3 {
					case "eat":
						fmt.Println(made_bird.Eat())
					case "move":
						fmt.Println(made_bird.Move())
					case "speak":
						fmt.Println(made_bird.Speak())
					default:
						fmt.Printf("%v is not a valid command.\n", command3)
					}
				case made_snake.name:
					switch command3 {
					case "eat":
						fmt.Println(made_snake.Eat())
					case "move":
						fmt.Println(made_snake.Move())
					case "speak":
						fmt.Println(made_snake.Speak())
					default:
						fmt.Printf("%v is not a valid command.\n", command3)
					}
				default:
					fmt.Println("There's no animal with that name.")
					continue
				}
			}

		} else {
			fmt.Printf("%v is not a valid command.\n", command1)
		}
	}
}
