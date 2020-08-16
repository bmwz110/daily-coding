package main

import (
	"fmt"
)

func main() {
	heights := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(heights))
}

func majorityElement(nums []int) int {
	major, count := 0, 0

	for _, num := range nums {
		if count == 0 {
			major = num
		}

		if major == num {
			count++
		} else {
			count--
		}
	}
	return major
}
