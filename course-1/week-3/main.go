package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	sl := make([]int, 0, 3)

	for {
		fmt.Printf("Enter 'X' to exit\nOR\nEnter an integer below\n\n\n")

		var input string
		_, err := fmt.Scan(&input)

		if input == "X" {
			break
		}

		integer, err := strconv.Atoi(input)

		if err != nil {
			fmt.Printf("Incorrect input type\n\n\n")
			continue
		}

		sl = append(sl, int(integer))

		sort.Ints(sl)
		fmt.Println(sl)

		fmt.Print("\n\n\n")
	}

	fmt.Printf("Goodbye!\n")
}
