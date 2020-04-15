// Copyright (c) 2020 Amit Levy
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package network

import (
	"testing"

	"gonum.org/v1/gonum/graph"
)

func TestCreateGraph(t *testing.T) {
	g := NewGraph()
	n1 := g.NewNode()
	g.AddNode(n1)
	n2 := g.NewNode()
	g.AddNode(n2)
	g.SetWeightedEdge(g.NewWeightedEdge(n1, n2, 5))
	if !g.HasEdgeFromTo(n1.ID(), n2.ID()) {
		t.Errorf("graph not created as expected")
	} else {
		t.Logf("graph %#v created!", g)
	}
}

func TestBFSFromNonExistSrc(t *testing.T) {
	g := NewGraph()
	if node := g.BFS(1, func(v graph.Node, depth int) bool { return true }); node != nil {
		t.Error("should not BFS when src node not in graph")
	}
}

func TestBFSEarlyStop(t *testing.T) {
	g := NewGraph()
	n1 := g.NewNode()
	g.AddNode(n1)
	n2 := g.NewNode()
	g.AddNode(n2)
	g.SetWeightedEdge(g.NewWeightedEdge(n1, n2, 0))
	sum := int64(0)
	early := g.BFS(0, func(v graph.Node, depth int) bool {
		sum += v.ID()
		return true
	})
	if early != n1 {
		t.Error("expected to stop early")
	}
	if sum != 0 {
		t.Error("did not iterate before early stopping")
	}
}

func TestDFSNonExistSrc(t *testing.T) {
	g := NewGraph()
	if node := g.DFS(1, func(v graph.Node) bool { return true }); node != nil {
		t.Error("should not DFS when src node not in graph")
	}
}

func TestDFS(t *testing.T) {
	g := NewGraph()
	g.AddNode(g.NewNode())
	if node := g.DFS(0, func(v graph.Node) bool { return false }); node != nil {
		t.Error("should DFS on one node")
	}
}
