package doublepointers

func merge(nums1 []int, m int, nums2 []int, n int) {
	// You must iterate numbers from the tail to the head!
	p1 := m - 1
	p2 := n - 1
	destP := m + n - 1
	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] < nums2[p2] {
			nums1[destP] = nums2[p2]
			p2--
		} else {
			nums1[destP] = nums1[p1]
			p1--
		}
		destP--
	}
	for p1 >= 0 {
		nums1[destP] = nums1[p1]
		p1--
		destP--
	}
	for p2 >= 0 {
		nums1[destP] = nums2[p2]
		p2--
		destP--
	}
}
