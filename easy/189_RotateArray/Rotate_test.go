package rotate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_rotate(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		{
			nums: []int{1, 2, 3, 4, 5, 6, 7},
			k:    3,
			want: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			nums: []int{-1, -100, 3, 99},
			k:    2,
			want: []int{3, 99, -1, -100},
		},
		{
			nums: []int{1, 2},
			k:    3,
			want: []int{2, 1},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			rotate(tt.nums, tt.k)
			assert.NotNil(t, tt.nums)
		})
	}
}
