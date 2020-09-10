package main

import (
	"fmt"
	"time"
)

// A race condition in software is a condition in which the program's
// outcome is based on the order or timing of non-deterministic events.
// The below program exemplifies a basic race condition.

func race(sleep bool) int {
	// A variable is added as storage for an integer (a counter).
	counter := 0

	go func() {
		// The value of the variable is updated within the secondary thread.
		counter++
	}()

	if sleep {
		// This time delay can change the outcome of the program.
		time.Sleep(time.Second * 1)
	}

	// The value of the variable is updated within the main thread.
	counter++

	// Although it's probable that with a time delay (`sleep`=true), the
	// returned value will be 2, it's not guaranteed.
	return counter
}

func main() {
	fmt.Println("With one second delay:", race(true))
	fmt.Println("Without one second delay:", race(false))
}
