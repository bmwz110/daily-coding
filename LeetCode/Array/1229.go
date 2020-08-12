package main

import (
	"fmt"
)

func main() {
	nums := []int{17, 18, 5, 4, 6, 1}
	fmt.Println(replaceElements(nums))
}

func replaceElements(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	ans := make([]int, len(arr)+1)
	ans[len(ans)-1] = -1

	for i := len(arr) - 1; i > 0; i-- {
		ans[i] = max(ans[i+1], arr[i])
	}

	return ans[1:]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
