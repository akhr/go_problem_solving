package graph

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
				u := Node{Val: StringerInt(edge[0])}
				v := Node{Val: StringerInt(edge[1])}
				g.AddEdge(u, v)
			}
			g.PrintGraph()
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewDirectedGraph()
			for _, edge := range tt.args.edges {
				u := Node{Val: StringerInt(edge[0])}
				v := Node{Val: StringerInt(edge[1])}
				g.AddEdge(u, v)
			}
			g.PrintGraph()
			g.TopoSort()
		})
	}
}

type StringerInt int

func (s StringerInt) String() string {
	return strconv.Itoa(int(s))
}
