package maxprofit

import "testing"

func Test_maxProfit(t *testing.T) {
	tests := []struct {
		prices []int
		want   int
	}{
		{
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   7,
		},
		{
			prices: []int{1, 2, 3, 4, 5},
			want:   4,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := maxProfit(tt.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
			t.Logf("MaxProfit :: %d\n", maxProfit(tt.prices))
		})
	}
}
