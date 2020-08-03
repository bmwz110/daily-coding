func merge(nums1 []int, m int, nums2 []int, n int) []int  {
	p1 := m - 1
	p2 := n - 1
	p := m + n - 1

    // 边界条件
	if p1 < 0 && p2 < 0 {
		return nil
	} else if p1 < 0 && p2 >= 0 {
		return nums2
	} else if p1 >= 0 && p2 < 0 {
		return nums1
	}

    // 从后往前比较，大的数放入nums1的末尾 
	for (p1 >= 0) && (p2 >= 0) {
		if nums1[p1] < nums2[p2] {
			nums1[p] = nums2[p2]
			p2--
		} else {
			nums1[p] = nums1[p1]
			p1--
		}
		p--
	}

	return nums1
}
