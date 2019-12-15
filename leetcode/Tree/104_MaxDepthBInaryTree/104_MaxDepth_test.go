package maxdepth104

import "testing"

func Test_maxDepth(t *testing.T) {
	tests := []struct {
		items []int
		want  int
	}{
		{
			items: []int{3, 9, 20, -1, -1, 15, 7},
			want:  3,
		},
		{
			items: []int{},
			want:  0,
		},
		{
			items: []int{1, 2},
			want:  2,
		},
	}
	for _, tt := range tests {
		root := constructTree(tt.items, 0)
		t.Run("", func(t *testing.T) {
			if got := maxDepth_Pre(root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
