package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	acceleration, err := ScanFloat64("acceleration")
	initialVelocity, err := ScanFloat64("initial velocity")
	initialDisplacement, err := ScanFloat64("initial displacement")

	if err != nil {
		fmt.Println("An incorrect type was entered.")
		return
	}

	displaceFn := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)

	var time float64
	fmt.Println("Enter time :")
	fmt.Scan(&time)

	fmt.Println(displaceFn(time))
}

// GenDisplaceFn returns a function which computes displacement as a function of time
func GenDisplaceFn(
	a float64,
	iv float64,
	id float64,
) func(t float64) float64 {
	fn := func(t float64) float64 {
		return (1/2)*a*(math.Pow(t, 2)) + iv*t + id
	}
	return fn
}

// ScanFloat64 returns a float64 entered by user
func ScanFloat64(name string) (float64, error) {
	var input string
	fmt.Println("Enter", name, ":")
	fmt.Scan(&input)
	return strconv.ParseFloat(input, 64)
}
