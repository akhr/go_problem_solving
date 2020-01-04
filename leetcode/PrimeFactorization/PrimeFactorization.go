package prime

import "math"

func getPrimeFactors(n int) []int {
	if n <= 0 {
		return []int{}
	}

	res := []int{}
	for n%2 == 0 {
		res = append(res, 2)
		n /= 2
	}

	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		for n%i == 0 {
			res = append(res, i)
			n /= i
		}
	}
	return res
}
