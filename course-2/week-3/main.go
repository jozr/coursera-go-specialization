package main

import (
	"errors"
	"fmt"
)

func main() {
	printInstructions()

	for {
		var animalInput string
		var actionInput string

		fmt.Print("> ")

		_, err := fmt.Scanln(&animalInput, &actionInput)
		if err != nil {
			if animalInput == "X" {
				break
			}

			fmt.Println("Unexpected input. Try again.")
			continue
		}

		animal, err := createAnimal(animalInput)
		if err != nil {
			fmt.Println(err)
			continue
		}

		answer, err := animal.getActionAnswer(actionInput)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(answer)
	}
}

func printInstructions() {
	fmt.Println("---------------------------------------------------")
	fmt.Println("Enter an animal and attribute separated by a space.")
	fmt.Println()
	fmt.Println("Animal options: cow, bird, snake")
	fmt.Println("Attribute options: eat, move, speak")
	fmt.Println()
	fmt.Println("Example:")
	fmt.Println("> bird eat")
	fmt.Println()
	fmt.Println("Enter 'X' to exit.")
	fmt.Println("---------------------------------------------------")
	fmt.Println()
}

func createAnimal(name string) (Animal, error) {
	var animal Animal
	var err error

	switch name {
	case "cow":
		animal = Animal{"grass", "walk", "moo"}
	case "bird":
		animal = Animal{"worms", "fly", "peep"}
	case "snake":
		animal = Animal{"mice", "slither", "hsss"}
	default:
		err = errors.New("The animal type was unrecognized. Please choose from: cow, bird, snake")
	}

	return animal, err
}

func (animal Animal) getActionAnswer(action string) (string, error) {
	var answer string
	var err error

	switch action {
	case "eat":
		answer = animal.Eat()
	case "move":
		answer = animal.Move()
	case "speak":
		answer = animal.Speak()
	default:
		err = errors.New("The animal action was unrecognized. Please choose from: eat, move, speak")
	}

	return answer, err
}

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (animal Animal) Eat() string {
	return animal.food
}

func (animal Animal) Move() string {
	return animal.locomotion
}

func (animal Animal) Speak() string {
	return animal.noise
}
