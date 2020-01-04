package graph

import (
	"fmt"
)

func (g *DirectedGraph) TopoSort() []Node {
	visited := map[Node]struct{}{}
	orderedSet := NewOrderedSet()
	for n, _ := range g.Nodes {
		g.topoSortHelper(n, visited, orderedSet)
	}

	res := []Node{}

	fmt.Println("Topo Sorted Order of Nodes :: ")
	for _, node := range orderedSet.GetItems() {
		fmt.Println(node.Val.String())
		res = append(res, node)
	}
	return res
}

func (g *DirectedGraph) topoSortHelper(node Node, visited map[Node]struct{}, orderedSet *OrderedSet) {
	fmt.Printf("TopoSortHelper - Start - %s\n", node.Val.String())
	if _, ok := visited[node]; !ok {
		visited[node] = struct{}{}
	}
	adjNodes := g.Edges[node]
	for adjNode, _ := range adjNodes {
		if _, ok := visited[adjNode]; !ok {
			g.topoSortHelper(adjNode, visited, orderedSet)
		}
	}
	fmt.Printf("TopoSortHelper - End - PushToList %s\n", node.Val.String())
	orderedSet.Add(node)
}

type OrderedSet struct {
	indexMap map[Node]int
	items    []Node
	length   int
}

func NewOrderedSet() *OrderedSet {
	return &OrderedSet{
		indexMap: map[Node]int{},
		items:    []Node{},
		length:   0,
	}
}

func (o *OrderedSet) Add(node Node) {
	if _, ok := o.indexMap[node]; !ok {
		o.indexMap[node] = o.length
		o.items = append(o.items, node)
		o.length++
	}
}

func (o *OrderedSet) GetItems() []Node {
	return o.items
}
