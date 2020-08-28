package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string

	fmt.Printf("Please enter string:\n")

	_, err := fmt.Scan(&str)

	if !found(str) || err != nil {
		fmt.Printf("Not Found!\n")
	} else {
		fmt.Printf("Found!\n")
	}
}

func found(str string) bool {
	strLength := len(str)
	lowerCaseStr := strings.ToLower(str)

	firstChar := string(lowerCaseStr[0])
	lastChar := string(lowerCaseStr[strLength-1])

	if firstChar == "i" && lastChar == "n" && strings.Contains(lowerCaseStr, "a") {
		return true
	}

	return false
}
