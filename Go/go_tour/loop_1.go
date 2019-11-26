package main

import (
	"fmt"
)

// Sqrt is a function
func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * x)
		fmt.Println(z)
	}
	return z
}

func main() {
	Sqrt(3)
}
