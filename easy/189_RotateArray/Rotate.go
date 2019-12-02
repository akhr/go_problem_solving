package rotate

func reverse(s []int, begin, end int) {
	if len(s) <= 1 || begin >= end {
		return
	}

	for i, j := begin, end; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rotate(nums []int, k int) {
	if k >= len(nums) {
		k %= len(nums)
	}
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}
