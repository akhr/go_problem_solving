package containsDuplicate217

import "testing"

func Test_containsDuplicate_sort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{
				nums: []int{1, 2, 3, 1},
			},
			want: true,
		},
		{
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: false,
		},
		{
			args: args{
				nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := containsDuplicate_sort(tt.args.nums); got != tt.want {
				t.Errorf("containsDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_containsDuplicate_map(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{
				nums: []int{1, 2, 3, 1},
			},
			want: true,
		},
		{
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: false,
		},
		{
			args: args{
				nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := containsDuplicate_map(tt.args.nums); got != tt.want {
				t.Errorf("containsDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
