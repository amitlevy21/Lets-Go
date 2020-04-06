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

// BFS breadth-first-search algorithm, starts from given start node
// invokes fn function when visiting every node, if fn return true
// the function stops and returns true
func (g *Graph) BFS(startNodeID int64, fn func(n graph.Node, depth int) bool) graph.Node {
	start := g.Node(startNodeID)
	if start == nil {
		return nil
	}
	BFS := traverse.BreadthFirst{}
	return BFS.Walk(g, start, fn)
}
