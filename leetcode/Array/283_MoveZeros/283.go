package movezeros283

//Two pointer solution
func moveZeroes(nums []int) {
	if len(nums) == 1 {
		return
	}
	lastNonZeroIndx := 0
	for _, num := range nums {
		if num != 0 {
			nums[lastNonZeroIndx] = num
			lastNonZeroIndx++
		}
	}

	for i := lastNonZeroIndx; i < len(nums); i++ {
		nums[i] = 0
	}
}
