package containsDuplicate217

import (
	"fmt"
	"sort"
)

func containsDuplicate_sort(nums []int) bool {
	sort.Ints(nums)
	i := 1
	for i < len(nums) {
		if nums[i-1] == nums[i] {
			return true
		}
		i++
	}
	return false
}

func containsDuplicate_map(nums []int) bool {
	seen := make(map[int]struct{}, len(nums))
	for _, n := range nums {
		if _, ok := seen[n]; ok {
			return true
		}
		seen[n] = struct{}{}
		fmt.Println(seen)
	}
	return false
}
