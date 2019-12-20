package mergesortedlists88

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_merge(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{
				nums1: []int{1, 2, 3, 0, 0, 0},
				m:     3,
				nums2: []int{2, 5, 6},
				n:     3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
			assert.Equal(t, []int{1, 2, 2, 3, 5, 6}, tt.args.nums1)
			t.Logf("Merged List :: %+v", tt.args.nums1)
		})
	}
}

func Test_merge_builtin(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{
				nums1: []int{1, 2, 3, 0, 0, 0},
				m:     3,
				nums2: []int{2, 5, 6},
				n:     3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge_buitin(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
			assert.Equal(t, []int{1, 2, 2, 3, 5, 6}, tt.args.nums1)
			t.Logf("Merged List :: %+v", tt.args.nums1)
		})
	}
}
