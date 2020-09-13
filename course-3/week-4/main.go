package main

import (
	"fmt"
	"sync"
)

var limit = 5
var dinner sync.WaitGroup

// Chopstick represents the eating utensil with
// mutex concurrency safety
type Chopstick struct{ sync.Mutex }

// Philosopher represents a numbered dinner guest
// with the ability to hold two chopsticks
type Philosopher struct {
	number                        int
	leftChopstick, rightChopstick *Chopstick
}

func (p Philosopher) eat() {
	for i := 0; i < 3; i++ {
		p.leftChopstick.Lock()
		p.rightChopstick.Lock()

		fmt.Println("starting to eat", p.number)
		fmt.Println("finishing eating", p.number)

		p.leftChopstick.Unlock()
		p.rightChopstick.Unlock()
	}

	dinner.Done()
}

func initChopsticks() []*Chopstick {
	chopsticks := make([]*Chopstick, limit)
	for i := range chopsticks {
		chopsticks[i] = new(Chopstick)
	}

	return chopsticks
}

func main() {
	chopsticks := initChopsticks()
	philosophers := make([]*Philosopher, limit)

	for i := range philosophers {
		number := i + 1
		philosophers[i] = &Philosopher{
			number,
			chopsticks[i],
			chopsticks[number%5]}

		dinner.Add(1)
		go philosophers[i].eat()
	}

	// Wait until everyone is finished eating
	dinner.Wait()
}
