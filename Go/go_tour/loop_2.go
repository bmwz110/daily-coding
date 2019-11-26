package main

import (
	"fmt"
	"math"
)

// Sqrt is a function
func Sqrt(x float64) float64 {
	z := 2.0
	num := 1
	for math.Abs(x-z*z) > 0.01 {
		z -= (z*z - x) / (2 * x)
		num++
		fmt.Println(z)
	}
	fmt.Println("The calculate number is: ", num)
	return z
}

func main() {
	Sqrt(3)
}
