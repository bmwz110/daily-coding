package main

import (
	"fmt"
)

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}

	fmt.Println(findDisappearedNumbers(nums))
}

func findDisappearedNumbers(nums []int) []int {
	arr := make([]int, len(nums)+1)
	arr2 := make([]int, 0)

	for _, v := range nums {
		arr[v]++
	}

	for i, v := range arr {
		if v == 0 {
			arr2 = append(arr2, i)
		}
	}

	return arr2[1:]
}
