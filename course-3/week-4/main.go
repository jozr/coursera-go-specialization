package main

import (
	"fmt"
	"sync"
	"time"
)

var limit = 5
var dinner sync.WaitGroup
var host = make(chan int, 2)

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
		host <- 1
		p.leftChopstick.Lock()
		p.rightChopstick.Lock()

		fmt.Println("starting to eat", p.number)
		time.Sleep(1 * time.Second)
		fmt.Println("finishing eating", p.number)

		p.leftChopstick.Unlock()
		p.rightChopstick.Unlock()
		<-host
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

func initPhilosophers(chopsticks []*Chopstick) []*Philosopher {
	philosophers := make([]*Philosopher, limit)
	for i := range philosophers {
		number := i + 1
		philosophers[i] = &Philosopher{
			number, chopsticks[i], chopsticks[number%5]}
	}
	return philosophers
}

func main() {
	chopsticks := initChopsticks()
	philosophers := initPhilosophers(chopsticks)

	for _, philosopher := range philosophers {
		dinner.Add(1)
		go philosopher.eat()
	}

	// Wait until everyone is finished eating
	dinner.Wait()
}
