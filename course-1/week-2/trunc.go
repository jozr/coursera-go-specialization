package main

import "fmt"

func main() {
	var f float64

	fmt.Printf("Please enter a float:\n")

	_, err := fmt.Scanln(&f)
	i := int64(f)

	if err != nil {
		fmt.Printf("You entered an invalid type.\n")
		return
	}

	if float64(i) == f {
		fmt.Printf("You entered an integer.\n")
		return
	}

	fmt.Printf("Integer: %+v\n", i)
}
