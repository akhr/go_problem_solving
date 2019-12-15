package plusone66

func plusOne(digits []int) []int {
	var n int = len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i] += 1
			return digits
		} else {
			digits[i] = 0
		}
	}
	var a = make([]int, n+1)
	a[0] = 1
	return a

}

// plusone.go
func PlusOne_2(digits []int) []int {
	// add one more slot to head for convenient
	digits = append([]int{0}, digits...)
	for i := len(digits) - 1; i >= 0; i-- {
		// find the first digit not 9 and plus one
		if digits[i] != 9 {
			digits[i] += 1
			break
		} else {
			// otherwise set to 0
			digits[i] = 0
		}
	}
	// if the first digit is 0 return the rest of slice
	if digits[0] == 0 {
		return digits[1:]
	}
	return digits
}
