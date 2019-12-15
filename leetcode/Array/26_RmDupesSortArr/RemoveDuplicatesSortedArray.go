package removedupes

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return 1
	}
	swapIdx := 1
	prev := nums[0]
	for _, num := range nums {
		if num != prev {
			prev = num
			nums[swapIdx] = num
			swapIdx++
		}
	}
	return swapIdx
}

func removeDuplicates_2(nums []int) int {
	if len(nums) <= 1 {
		return 1
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
