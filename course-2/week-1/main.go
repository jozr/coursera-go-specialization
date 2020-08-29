package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var input string

	fmt.Println("Enter a up to 10 sequential integers without spaces:")
	fmt.Scan(&input)

	integers, err := ParseUserInput(input)

	if err != nil {
		fmt.Println("Incorrect input type. Please enter integers.")
		return
	}

	if len(integers) > 10 {
		fmt.Println("There are more than 10 integers.")
		return
	}

	fmt.Println(BubbleSort(integers))
}

// ParseUserInput parses the user input into a slice of integers
func ParseUserInput(input string) ([]int, error) {
	numbers := strings.Split(input, "")
	integers := make([]int, 0, len(numbers))

	for _, number := range numbers {
		integer, err := strconv.Atoi(number)

		if err != nil {
			return nil, err
		}

		integers = append(integers, integer)
	}

	return integers, nil
}

// BubbleSort sorts a series of integers from least to greatest
func BubbleSort(integers []int) []int {
	loopLen := len(integers) - 1

	for i := 0; i < loopLen; i++ {
		swapped := false

		for j := 0; j < loopLen; j++ {
			if integers[j] > integers[j+1] {
				Swap(integers, j)
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}

	return integers
}

// Swap takes a slice and index and sorts the current and next value
func Swap(integers []int, i int) {
	integers[i], integers[i+1] = integers[i+1], integers[i]
}
