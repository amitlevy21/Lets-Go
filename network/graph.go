// Copyright (c) 2020 Amit Levy
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package network

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
