package main

import (
	"fmt"
	"strconv"
)

func main() {
	integers := make([]int, 0, 10)

	for i := 0; i < 10; i++ {
		var input string

		fmt.Println("Enter an integer:")
		fmt.Scan(&input)

		integer, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Incorrect input type. Please enter integers.")
			return
		}

		integers = append(integers, integer)
	}

	fmt.Println(BubbleSort(integers))
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
