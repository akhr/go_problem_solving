package singlenumber136

func singleNumber(nums []int) int {
	var result int
	for _, num := range nums {
		result = result ^ num
	}
	return result
}
