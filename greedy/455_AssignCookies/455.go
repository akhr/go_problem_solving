package assigncookies455

import "sort"

func findContentChildren_2(g []int, s []int) int {
	if len(g) == 0 || len(s) == 0 {
		return 0
	}
	maxChildren := 0
	sort.Sort(sort.IntSlice(s))
	sort.Sort(sort.IntSlice(g))
	for _, child := range g {
		for i, cookie := range s {
			if child <= cookie {
				maxChildren++
				s = s[i+1:] // remove all smaller cookies which cannot be assinged to any child.
				break
			}
		}
	}
	return maxChildren
}

func findContentChildren(g []int, s []int) int {
	maxChildren := 0
	sort.Sort(sort.IntSlice(s))
	sort.Sort(sort.IntSlice(g))
	chld := 0
	cookie := 0
	for chld < len(g) && cookie < len(s) {
		if g[chld] <= s[cookie] {
			maxChildren++
			chld++
		}
		cookie++
	}
	return maxChildren
}
