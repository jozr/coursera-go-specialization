package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	printInstructions()

	integers := make([]int, 0)

	for {
		var input string
		_, err := fmt.Scan(&input)
		printLine()

		if input == "X" {
			printLine()
			return
		}

		if input == "S" {
			fmt.Println("Unsorted integers:", integers)
			sortedIntegers := Sort(integers)
			fmt.Println("Sorted integers:", sortedIntegers)
			printLine()
			return
		}

		integer, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Incorrect input type. Try again.")
			printLine()
			continue
		}

		integers = append(integers, int(integer))
	}
}

var amount = 4

// Sort array of integers from least to greatest
func Sort(integers []int) []int {
	channel := make(chan []int, amount)
	sorted := make([]int, 0)

	// sort the 4 slices each in a routine
	for index, partition := range part(integers) {
		go func(i int, p []int) {
			fmt.Println("Routine", i+1, "will sort:", p)

			sort.Ints(p)
			channel <- p
		}(index, partition)
	}

	// wait for result of each routine
	for i := 0; i < amount; i++ {
		sorted = append(sorted, <-channel...)
	}

	sort.Ints(sorted) // final sort on the entire slice
	return sorted
}

func part(slice []int) (partitions [][]int) {
	size := sizePartitions(len(slice))

	for size < len(slice) {
		slice, partitions = slice[size:], append(partitions, slice[0:size:size])
	}

	return append(partitions, slice)
}

func sizePartitions(length int) int {
	float := float32(length) / float32(amount)

	return int(float + 0.5)
}

func printInstructions() {
	fmt.Println("Enter as many integers as you like.")
	fmt.Println("When you're finished, enter 'S' to sort the integers.")
	fmt.Println("Enter 'X' to exit.")
	printLine()
}

func printLine() {
	fmt.Println("-----------------------------------------------------")
}
