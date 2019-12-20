package intersectionofarrays350

func intersect(nums1 []int, nums2 []int) []int {
	var shorter, longer []int
	if len(nums1) < len(nums2) {
		shorter = nums1
		longer = nums2
	} else {
		shorter = nums2
		longer = nums1
	}

	m := map[int]int{}
	for _, num := range longer {
		m[num]++
	}

	res := []int{}
	for _, num := range shorter {
		if m[num] > 0 {
			m[num]--
			res = append(res, num)
		}
	}
	return res
}
