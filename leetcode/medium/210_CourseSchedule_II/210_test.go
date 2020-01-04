package coursescheduleII210

import (
	"strconv"
	"testing"
)

func TestDirectedGraph_AddEdge(t *testing.T) {
	type args struct {
		edges [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{
				edges: [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewDirectedGraph()
			for _, edge := range tt.args.edges {
				u := Node{Val: edge[0]}
				v := Node{Val: edge[1]}
				g.AddEdge(u, v)
			}
		})
	}
}

func TestDirectedGraph_TopoSort(t *testing.T) {
	type args struct {
		edges [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{
				edges: [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			},
		},
		{
			args: args{
				edges: [][]int{{}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewDirectedGraph()
			for _, edge := range tt.args.edges {
				u := Node{Val: edge[0]}
				v := Node{Val: edge[1]}
				g.AddEdge(u, v)
			}
			g.TopoSort()
		})
	}
}

type StringerInt int

func (s StringerInt) String() string {
	return strconv.Itoa(int(s))
}

func Test_findOrder(t *testing.T) {
	type args struct {
		num   int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		/* {
			args: args{
				num:   4,
				edges: [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			},
		}, */
		{
			args: args{
				num:   1,
				edges: [][]int{{}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := findOrder(tt.args.num, tt.args.edges)
			t.Logf("Res :: % +v", res)
		})
	}
}
