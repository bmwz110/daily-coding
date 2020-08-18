package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 3, 3, 3, 3, 3, 312}
	fmt.Println(moveZeroes(nums))
}

func moveZeroes(nums []int) int {
	num, vote := 0, 0
	for _, v := range nums {
		if vote == 0 {
			num = v
		}

		if v != num {
			vote--
		} else {
			vote++
		}
	}
	return num
}
