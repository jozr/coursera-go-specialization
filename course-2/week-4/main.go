package main

import (
	"errors"
	"fmt"
)

func main() {
	printInstructions()
	animals := make(map[string]Animal)

	for {
		var command string
		var name string
		var info string

		fmt.Print("> ")

		_, err := fmt.Scanln(&command, &name, &info)
		if err != nil {
			if command == "X" {
				break
			}

			fmt.Println("Unexpected input. Try again.")
			continue
		}

		if command == "newanimal" {
			animal, err := createAnimal(name, info)
			if err != nil {
				fmt.Println(err)
				continue
			}

			animals[name] = animal

			fmt.Println(name, "was created.")
			continue
		}

		if command == "query" {
			queryAnimal := animals[name]
			if queryAnimal == nil {
				fmt.Println("The animal does not exist yet.")
				continue
			}

			printActionAnswer(animals[name], info)
		}
	}
}

func printInstructions() {
	fmt.Println("--------------------------------------")
	fmt.Println("Enter a command separated by a spaces.")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("> newanimal fatima bird")
	fmt.Println("> query fatima move")
	fmt.Println()
	fmt.Println("Enter 'X' to exit.")
	fmt.Println("--------------------------------------")
	fmt.Println()
}

func createAnimal(name string, category string) (Animal, error) {
	var animal Animal
	var err error

	switch category {
	case "cow":
		animal = Cow{}
	case "bird":
		animal = Bird{}
	case "snake":
		animal = Snake{}
	default:
		err = errors.New("The animal type was unrecognized. Please choose from: cow, bird, snake")
	}

	return animal, err
}

func printActionAnswer(animal Animal, action string) {
	switch action {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		fmt.Println("The animal action was unrecognized. Please choose from: eat, move, speak")
	}
}

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{ name string }

func (Cow) Eat() {
	fmt.Println("grass")
}

func (Cow) Move() {
	fmt.Println("walk")
}

func (Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct{ name string }

func (Bird) Eat() {
	fmt.Println("worms")
}

func (Bird) Move() {
	fmt.Println("fly")
}

func (Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct{ name string }

func (Snake) Eat() {
	fmt.Println("mice")
}

func (Snake) Move() {
	fmt.Println("slither")
}

func (Snake) Speak() {
	fmt.Println("hsss")
}
