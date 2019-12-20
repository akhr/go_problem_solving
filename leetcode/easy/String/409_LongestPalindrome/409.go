package longestpalindrome409

import (
	"fmt"
)

func longestPalindrome(s string) int {
	runes := []rune(s)
	m := make(map[rune]int, len(s))
	for _, r := range runes {
		if cnt, ok := m[r]; ok {
			m[r] = cnt + 1
		} else {
			m[r] = 1
		}
	}
	fmt.Printf("map - % +v", m)

	len, odd := 0, 0
	for _, n := range m {
		if n%2 == 0 {
			len += n
		} else {
			len += (n / 2) * 2
			odd = 1
		}
	}
	return len + odd
}
