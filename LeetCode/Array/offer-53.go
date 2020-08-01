https://leetcode-cn.com/problems/que-shi-de-shu-zi-lcof/

// Solution 1
func missingNumber(nums []int) int {
	left := 0
	right := len(nums)
	for left < right {
		mid := (left + right) >> 1   // 二进制向右移 n 位，等同于除以 2 的 n 次方
		if nums[mid] != mid {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}



// Solution 2
func missingNumber(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	// 不缺失情况下的和 total 减去 缺失情况下的和 sum，即为缺失值
	total := (1 + nums[len(nums)-1]) * nums[len(nums)-1] / 2

	return total - sum
}



