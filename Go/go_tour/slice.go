package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	// 添加一个空切片
	s = append(s, 0)
	printSlice(s)

	// 这个切片会按需增长
	s = append(s, 1)
	printSlice(s)

	// 可以一次性添加多个元素
	// Slice扩容机制:
	// 1. append单个元素，或者append少量的多个元素，这里的少量指double之后的容量能容纳，这样就会走以下扩容流程，
	// 不足1024，双倍扩容，超过1024的，1.25倍扩容;
	// 2. 若是append多个元素，且double后的容量不能容纳，直接使用预估的容量;
	// 以上两个分支得到新容量后，均需要根据slice的类型size，算出新的容量所需的内存情况capmem，
	// 然后再进行capmem向上取整，得到新的所需内存，除上类型size，得到真正的最终容量,作为新的slice的容量.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
