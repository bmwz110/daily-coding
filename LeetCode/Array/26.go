func removeDuplicates(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

    // 双指针 
	p, q := 0, 1
	for q < len(nums) {
		if nums[p] != nums[q] {
			nums[p+1] = nums[q]
			p++
		}
		q++
	}
	return p + 1 
}
