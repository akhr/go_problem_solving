package issubsequence392

import "bytes"

func isSubsequence(s, t string) bool {
	x := []byte(t)
	for i := 0; i < len(s); i++ {
		indx := bytes.IndexByte(x, s[i])
		if indx == -1 {
			return false
		}
		x = x[indx+1:]
	}
	return true
}

/* func isSubsequence(s string, t string) bool {
	m := map[byte][]int{}
	for i, char := range []byte(t) {
		if _, ok := m[char]; !ok {
			m[char] = []int{i}
		} else {
			m[char] = append(m[char], i)
		}
	}

	prev := -1
	for _, char := range []byte(s) {
		if _, ok := m[char]; ok {
			if ok, indx := isCharAvailable(m[char], prev); ok {
				prev = indx
				continue
			}
		}
		return false
	}
	return true
}

func isCharAvailable(indxArr []int, prev int) (bool, int) {
	for _, indx := range indxArr {
		if indx > prev {
			return true, indx
		}
	}
	return false, -1
} */
