package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 18
	twoSum(nums, target)
}

func twoSum(nums []int, target int) []int {
	tempMap := map[int]int{}
	for i, v := range nums {
		g := target - v

		if j, ok := tempMap[g]; ok {
			return []int{j, i}
		}
		tempMap[v] = i
	}
	return []int{}
}
