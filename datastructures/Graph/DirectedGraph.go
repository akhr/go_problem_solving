package graph

import "fmt"

type Node struct {
	Val fmt.Stringer
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

func (g *DirectedGraph) PrintAllAdjNodes(node Node) {
	if _, ok := g.Edges[node]; !ok {
		fmt.Printf("No adjacent nodes exist for Node %s\n", node.Val.String())
		return
	}
	adjNodes := g.Edges[node]
	for adjNode, _ := range adjNodes {
		fmt.Printf("%s -> %s\n", node.Val.String(), adjNode.Val.String())
	}
}

func (g *DirectedGraph) PrintGraph() {
	for node, _ := range g.Nodes {
		g.PrintAllAdjNodes(node)
	}
}
