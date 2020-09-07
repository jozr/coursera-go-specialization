package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("With one second delay:", race(true))
	fmt.Println("Without one second delay:", race(false))
}

func race(sleep bool) int {
	counter := 0

	go func() {
		counter++
	}()

	if sleep {
		time.Sleep(time.Second * 1)
	}

	counter++

	return counter
}
