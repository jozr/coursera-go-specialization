package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	var fileName string
	fmt.Println("Name of file:")
	fmt.Scan(&fileName)

	bytesRead, _ := ioutil.ReadFile(fileName)
	content := string(bytesRead)
	trimmedContent := strings.TrimSuffix(content, "\n")
	rawNames := strings.Split(trimmedContent, "\n")

	type person struct {
		fname string
		lname string
	}
	persons := make([]person, 0, len(rawNames))

	for _, rawPerson := range rawNames {
		names := strings.Split(rawPerson, " ")
		persons = append(persons, person{names[0], names[1]})
	}

	for _, person := range persons {
		fmt.Println(person)
	}
}
