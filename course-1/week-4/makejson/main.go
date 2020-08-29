package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	person := make(map[string]string)

	var name string
	fmt.Println("Name:")
	fmt.Scan(&name)
	person["name"] = name

	var address string
	fmt.Println("Address:")
	fmt.Scan(&address)
	person["address"] = address

	personJSON, _ := json.Marshal(person)
	fmt.Println(string(personJSON))
}
