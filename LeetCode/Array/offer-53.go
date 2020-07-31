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
    n:=len(nums)
    count:=(n+n*n) >> 1
    for _,v:=range nums{
        count-=v
    }
    return count
}



// Solution 3
func missingNumber(nums []int) int {
    var count=len(nums)
    for i,v:=range nums{
        count^=i^v
    }
    return count
}
