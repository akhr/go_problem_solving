package removedupes

import "testing"

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{1, 1, 2},
			want: 2,
		},
		{
			nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := removeDuplicates_2(tt.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
			t.Logf("nums :: %#v", tt.nums)
		})
	}
}
