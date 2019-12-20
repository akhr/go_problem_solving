package singlenumber136

import "testing"

func Test_singleNumber(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{2, 2, 1},
			want: 1,
		},
		{
			nums: []int{4, 1, 2, 1, 2},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := singleNumber(tt.nums); got != tt.want {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
