package coursescheduleII210

import (
	"errors"
	"fmt"
)

func findOrder(numCourses int, prerequisites [][]int) []int {
	if len(prerequisites) == 0 || len(prerequisites[0]) == 0 {
		res := []int{}
		for i := 0; numCourses > 0; numCourses-- {
			res = append(res, i)
			i++
		}
		return res
	}

	g := NewDirectedGraph()
	for _, edge := range prerequisites {
		u := Node{Val: edge[0]}
		v := Node{Val: edge[1]}
		g.AddEdge(u, v)
	}
	sortedList := g.TopoSort()
	return sortedList
}

type Node struct {
	Val int
}

type DirectedGraph struct {
	Nodes map[Node]struct{}
	Edges map[Node]map[Node]struct{}
}

func NewDirectedGraph() *DirectedGraph {
	g := &DirectedGraph{
		Nodes: map[Node]struct{}{},
		Edges: map[Node]map[Node]struct{}{},
	}
	return g
}

func (g *DirectedGraph) AddNode(node Node) {
	if _, ok := g.Nodes[node]; !ok {
		g.Nodes[node] = struct{}{}
	}
}

func (g *DirectedGraph) AddEdge(u, v Node) {
	if _, ok := g.Nodes[u]; !ok {
		g.Nodes[u] = struct{}{}
	}
	if _, ok := g.Nodes[v]; !ok {
		g.Nodes[v] = struct{}{}
	}
	if _, ok := g.Edges[u]; !ok {
		g.Edges[u] = map[Node]struct{}{}
	}
	g.Edges[u][v] = struct{}{}
}

func (g *DirectedGraph) TopoSort() []int {
	visited := map[Node]struct{}{}
	orderedSet := NewOrderedSet()
	for n, _ := range g.Nodes {
		g.topoSortHelper(n, visited, orderedSet)
	}

	res := []int{}
	fmt.Println("Topo Sorted Order of Nodes :: ")
	for _, node := range orderedSet.items {
		fmt.Println(node.Val)
		res = append(res, node.Val)
	}
	return res
}

func (g *DirectedGraph) topoSortHelper(node Node, visited map[Node]struct{}, orderedSet *OrderedSet) {
	if _, ok := visited[node]; !ok {
		visited[node] = struct{}{}
	}
	adjNodes := g.Edges[node]
	for adjNode, _ := range adjNodes {
		if _, ok := visited[adjNode]; !ok {
			g.topoSortHelper(adjNode, visited, orderedSet)
		}
	}
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

func (o *OrderedSet) Add(node Node) error {
	if _, ok := o.indexMap[node]; ok {
		return errors.New("Cyclic Dependency")
	}
	o.indexMap[node] = o.length
	o.items = append(o.items, node)
	o.length++
	return nil
}
