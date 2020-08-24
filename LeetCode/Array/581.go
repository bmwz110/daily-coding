package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{2, 6, 4, 8, 10, 9, 15}

	fmt.Println(findUnsortedSubarray(nums))
}

func findUnsortedSubarray(nums []int) int {
	sortedNums := make([]int, len(nums))
	copy(sortedNums, nums) // copy复制为值复制, =为引用复制
	sort.Slice(sortedNums, func(i, j int) bool {
		return sortedNums[i] < sortedNums[j]
	})

	result := 0

	start, end := len(nums), 0
	for i, v := range sortedNums {
		if v != nums[i] {
			start = min(start, i)
			end = max(end, i)
		}
	}

	if end-start >= 0 {
		result = end - start + 1
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
