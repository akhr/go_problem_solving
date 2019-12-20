package firstbadversion278

import (
	"fmt"
	"math"
)

/* The isBadVersion API is defined in the parent class VersionControl.
   boolean isBadVersion(int version); */

/* func isBadVersion(v int) bool {
	if v == 4 {
		return true
	}
	return false
} */

var isBadVersion func(int) bool

type firstBad struct {
	version        int
	isBadCallCount int
}

func firstBadVersion(n int, isBadV func(int) bool) int {
	isBadVersion = isBadV
	min := 1
	max := n
	firstBad := &firstBad{version: math.MaxInt64}
	helper(min, max, firstBad)
	fmt.Printf("# of calls to isBadVersion() :: %d\n", firstBad.isBadCallCount)
	return firstBad.version
}

func helper(min, max int, firstBad *firstBad) {
	if min > max {
		return
	}
	mid := (min + max) / 2
	firstBad.isBadCallCount++ // incr func call counter
	if isBadVersion(mid) {
		firstBad.version = int(math.Min(float64(firstBad.version), float64(mid)))
		helper(min, mid-1, firstBad)
	} else {
		helper(mid+1, max, firstBad)
	}
}
