package twosum

func twoSum(nums []int, target int) []int {
	indexMap := map[int]int{}
	for i, v := range nums {
		if j, ok := indexMap[target-v]; ok {
			return []int{j, i}
		}
		indexMap[v] = i
	}
	return []int{0, 0}
}
