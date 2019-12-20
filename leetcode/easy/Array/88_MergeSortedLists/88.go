package mergesortedlists88

import "sort"

func merge_buitin(nums1 []int, m int, nums2 []int, n int) {
	nums1 = append(nums1[0:m], nums2...)
	sort.Sort(sort.IntSlice(nums1))
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	p1, p2, p := m-1, n-1, len(nums1)-1
	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}
	for p1 >= 0 {
		nums1[p] = nums1[p1]
		p1--
		p--
	}

	for p2 >= 0 {
		nums1[p] = nums2[p2]
		p2--
		p--
	}
}
