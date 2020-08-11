package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 4, 5, 6}
	target := 4
	fmt.Println(searchInsert(nums, target))
}

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}
