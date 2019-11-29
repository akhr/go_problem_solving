package factorial

import "testing"

func Test_fact(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{n: 1},
			want: 1,
		},
		{
			args: args{n: 2},
			want: 2,
		},
		{
			args: args{n: 3},
			want: 6,
		},
		{
			args: args{n: 4},
			want: 24,
		},
		{
			args: args{n: 5},
			want: 120,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fact(tt.args.n); got != tt.want {
				t.Errorf("fact() = %v, want %v", got, tt.want)
			}
		})
	}
}
