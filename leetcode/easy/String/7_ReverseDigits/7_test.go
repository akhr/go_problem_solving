package reversedigits7

import "testing"

func Test_reverse(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want int
	}{
		{
			x:    123,
			want: 321,
		},
		{
			x:    -123,
			want: -321,
		},
		{
			x:    12,
			want: 21,
		},
		{
			x:    1534236469,
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
