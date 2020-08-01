package main

import "fmt"

func main() {
	heights := []int{1, 2, 1, 5, 6}
	fmt.Println(findLengthOfLCIS(heights))
}

func findLengthOfLCIS(nums []int) int {
	ans, anchor := 0, 0
	for i, _ := range nums {
		// 当递增规律中断时，重新定义 anchor
		if (i > 0) && nums[i-1] >= nums[i] {
			anchor = i
		}

		// i-anchor+1 为当前 递增序列长度
		// ans 为当前最长递增序列长度
		if ans <= i-anchor+1 {
			ans = i - anchor + 1
		}
	}
	return ans
}
