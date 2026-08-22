func merge(nums1 []int, m int, nums2 []int, n int) {
	p1,p2, k := m-1,n-1,m+n-1

	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[k] = nums1[p1]
			p1--
		} else {
			nums1[k] = nums2[p2]
			p2--
		}
		k--
	}

	for p2 >= 0 {
		nums1[k] = nums2[p2]
		p2--
		k--
	}
}
