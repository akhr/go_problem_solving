package factorial

func fact(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return n * fact(n-1)
}
