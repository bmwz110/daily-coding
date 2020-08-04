func merge(nums1 []int, m int, nums2 []int, n int) []int  {
	p1 := m - 1
	p2 := n - 1
	len := m + n - 1

	for p2 >= 0 {
		if p1 >= 0 && nums1[p1] > nums2[p2] {
			nums1[len] = nums1[p1]
			len--
			p1--
		} else {
			nums1[len] = nums2[p2]
			len--
			p2--
		}
	}

	return nums1
}
