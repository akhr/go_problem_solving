package firstbadversion278

import "testing"

func Test_firstBadVersion(t *testing.T) {
	type args struct {
		n     int
		isBad func(int) bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				n:     5,
				isBad: func(n int) bool { return n >= 4 },
			},
			want: 4,
		},
		{
			args: args{
				n:     8,
				isBad: func(n int) bool { return n >= 3 },
			},
			want: 3,
		},
		{
			args: args{
				n:     100,
				isBad: func(n int) bool { return n >= 93 },
			},
			want: 93,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstBadVersion(tt.args.n, tt.args.isBad); got != tt.want {
				t.Errorf("firstBadVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
