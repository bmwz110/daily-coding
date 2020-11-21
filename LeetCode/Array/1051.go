// 对数组进行排序，并返回与原数组的不同项个数

package main  

import "fmt"

func main() {
	heights := []int{5, 1, 2, 3, 4}
	fmt.Println(heightChecker(heights))
}

func heightChecker(heights []int) int {
	arr := [101]int{}
	for _, height := range heights { 
		arr[height]++ 
	} 

	count := 0
	for i, j := 1, 0; i < len(arr); i++ {
		for arr[i] > 0 {
			if heights[j] != i {
				count++
			}
			j++
			arr[i]--
		}
	}

	return count;
} 

