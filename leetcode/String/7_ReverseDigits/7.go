package reversedigits7

import "math"

func reverse(x int) int {
	isNeg := false
	if x < 0 {
		isNeg = true
		x = x * (-1)
	}
	rev := 0
	for x > 0 {
		rev = rev*10 + x%10
		x = x / 10
		if rev > math.MaxInt32 {
			return 0
		}
	}

	if isNeg {
		return rev * (-1)
	}
	return rev
}
