package main

import (
	"math"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/traverse"
)

// Graph is always a weighted directed graph
type Graph struct {
	*simple.WeightedDirectedGraph
}

// NewGraph must be used to initialize a graph.
// self pointing edges (loops) are not allowed and considered with infinite
// weight. The weight of a non existent edge is infinite.
func NewGraph() *Graph {
	loopWeight := math.Inf(1)
	absentWeight := math.Inf(1)
	return &Graph{simple.NewWeightedDirectedGraph(loopWeight, absentWeight)}
}

// BFS breadth-first-search algorithm, starts from given start node.
// invokes fn function when visiting every node, stops early if until returns true.
func (g *Graph) BFS(startNodeID int64, until func(n graph.Node, depth int) bool) graph.Node {
	start := g.Node(startNodeID)
	if start == nil {
		return nil
	}
	BFS := traverse.BreadthFirst{}
	return BFS.Walk(g, start, until)
}

// DFS depth-first-search algorithm, starts from given start node.
// invokes fn function when visiting every node, stops early if until returns true.
func (g *Graph) DFS(startNodeID int64, until func(n graph.Node) bool) graph.Node {
	start := g.Node(startNodeID)
	if start == nil {
		return nil
	}
	DFS := traverse.DepthFirst{}
	return DFS.Walk(g, start, until)
}

// Equal check if g and o are equal, meaning, they have the same nodes and edges.
// See EqualNodes and EqualEdges for more info.
func (g *Graph) Equal(o *Graph) bool {
	if g == o {
		return true
	}
	if !g.EqualNodes(o) {
		return false
	}
	if !g.EqualEdges(o) {
		return false
	}
	return true
}

// EqualNodes checks if g and o has the same nodes IDs.
func (g *Graph) EqualNodes(o *Graph) bool {
	gn := g.Nodes()
	on := o.Nodes()
	for gn.Next() {
		n1 := gn.Node().ID()
		if !on.Next() {
			return false
		}
		if n1 != on.Node().ID() {
			return false
		}
	}
	if on.Next() {
		return false
	}
	return true
}

// EqualEdges checks if g and o have the same pointing edges with same weights.
func (g *Graph) EqualEdges(o *Graph) bool {
	ge := g.WeightedEdges()
	oe := o.WeightedEdges()
	for ge.Next() {
		e1 := ge.WeightedEdge()
		if !oe.Next() {
			return false
		}
		e2 := oe.WeightedEdge()
		if e1.From().ID() != e2.From().ID() {
			return false
		}
		if e1.To().ID() != e2.To().ID() {
			return false
		}
		if e1.Weight() != e2.Weight() {
			return false
		}
	}
	if oe.Next() {
		return false
	}
	return true
}
