package movezeros283

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
		want []int
	}
	tests := []struct {
		args args
	}{
		{
			args{
				nums: []int{0, 1, 0, 3, 12},
				want: []int{1, 3, 12, 0, 0},
			},
		},
		{
			args{
				nums: []int{1},
				want: []int{1},
			},
		},
		{
			args{
				nums: []int{1, 0},
				want: []int{1, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			moveZeroes(tt.args.nums)
			t.Logf("Output-->%v", tt.args.nums)
			assert.Equal(t, tt.args.want, tt.args.nums)
		})
	}
}
