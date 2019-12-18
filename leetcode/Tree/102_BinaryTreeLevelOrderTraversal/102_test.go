package binarytreeleveltraversal102

import (
	"reflect"
	"testing"

	tree "github.com/go_problem_solving/leetcode/Tree"
)

func Test_levelOrder(t *testing.T) {
	tests := []struct {
		items []int
		want  [][]int
	}{
		{
			items: []int{3, 9, 20, -1, -1, 15, 7},
			want:  [][]int{{3}, {9, 20}, {15, 7}},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			root := tree.ConstructIntValTree(tt.items, 0)
			if got := levelOrder(root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_levelOrder_list_pkg(t *testing.T) {
	tests := []struct {
		items []int
		want  [][]int
	}{
		{
			items: []int{3, 9, 20, -1, -1, 15, 7},
			want:  [][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			items: []int{},
			want:  [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			root := tree.ConstructIntValTree(tt.items, 0)
			if got := levelOrder_list_pkg(root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
