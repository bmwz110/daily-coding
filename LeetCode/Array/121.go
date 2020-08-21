package main

import (
	"fmt"
)

func main() {
	nums := []int{7, 1, 5, 3, 6, 4}

	fmt.Println(maxProfit(nums))
}

func maxProfit(nums []int) int {
	// ^unit(0)在32位系统上返回 0xFFFFFFFF(即2^32)
	// 在64位系统上返回 0xFFFFFFFFFFFFFFFF (即2^64)
	// 可以利用 32 << (^uint(0) >> 63) 来得到系统位数
	minprice, bonus := int(^uint(0)>>1), 0

	for _, v := range nums {
		bonus = Max(v-minprice, bonus)
		minprice = Min(v, minprice)
	}

	return bonus
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
