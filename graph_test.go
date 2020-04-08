package main

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
		t.Logf("graph %v created!", g)
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

func TestEqualSelf(t *testing.T) {
	g1 := NewGraph()
	if !g1.Equal(g1) {
		t.Errorf("empty graphs should equal")
	}
}

func TestEqualEmpty(t *testing.T) {
	g1 := NewGraph()
	g2 := NewGraph()
	if !g1.Equal(g2) || !g2.Equal(g1) {
		t.Errorf("empty graphs should equal")
	}
}

func TestNotEqualNodes(t *testing.T) {
	g1 := NewGraph()
	g2 := NewGraph()
	g1.AddNode(g1.NewNode())
	if g1.Equal(g2) || g2.Equal(g1) {
		t.Errorf("graphs should not equal")
	}
}

func TestEqualNodes(t *testing.T) {
	g1 := NewGraph()
	g2 := NewGraph()
	g1.AddNode(g1.NewNode())
	g2.AddNode(g2.NewNode())
	if !g1.Equal(g2) || !g2.Equal(g1) {
		t.Errorf("graphs should equal")
	}
}

func TestNotEqualEdges(t *testing.T) {
	g1 := NewGraph()
	g2 := NewGraph()
	n1 := g1.NewNode()
	g1.AddNode(n1)
	n2 := g1.NewNode()
	g1.AddNode(n2)
	n1 = g2.NewNode()
	g2.AddNode(n1)
	n2 = g2.NewNode()
	g2.AddNode(n2)
	g1.SetWeightedEdge(g1.NewWeightedEdge(n1, n2, 0))
	if g1.Equal(g2) || g2.Equal(g1) {
		t.Errorf("g1 should not equal g2")
	}
}

func TestEqualEdgesToDiffNodes(t *testing.T) {
	g1 := NewGraph()
	g2 := NewGraph()
	n1 := g1.NewNode()
	g1.AddNode(n1)
	n2 := g1.NewNode()
	g1.AddNode(n2)
	g1.SetWeightedEdge(g1.NewWeightedEdge(n1, n2, 0))
	n1 = g2.NewNode()
	g2.AddNode(n1)
	n2 = g2.NewNode()
	g2.AddNode(n2)
	g2.SetWeightedEdge(g2.NewWeightedEdge(n2, n1, 0))
	if g1.Equal(g2) || g2.Equal(g1) {
		t.Errorf("graphs should not equal")
	}
}


func TestEqualEdgesToDiffNodes2(t *testing.T) {
	g1 := NewGraph()
	g2 := NewGraph()
	n1 := g1.NewNode()
	g1.AddNode(n1)
	n2 := g1.NewNode()
	g1.AddNode(n2)
	n3 := g1.NewNode()
	g1.AddNode(n3)
	g1.SetWeightedEdge(g1.NewWeightedEdge(n1, n2, 0))
	n1 = g2.NewNode()
	g2.AddNode(n1)
	n2 = g2.NewNode()
	g2.AddNode(n2)
	n3 = g2.NewNode()
	g2.AddNode(n3)
	g2.SetWeightedEdge(g2.NewWeightedEdge(n1, n3, 0))
	if g1.Equal(g2) || g2.Equal(g1) {
		t.Errorf("graphs should not equal")
	}
}

func TestEqualEdgesDiffWeight(t *testing.T) {
	g1 := NewGraph()
	g2 := NewGraph()
	n1 := g1.NewNode()
	g1.AddNode(n1)
	n2 := g1.NewNode()
	g1.AddNode(n2)
	g1.SetWeightedEdge(g1.NewWeightedEdge(n1, n2, 0))
	n1 = g2.NewNode()
	g2.AddNode(n1)
	n2 = g2.NewNode()
	g2.AddNode(n2)
	g2.SetWeightedEdge(g2.NewWeightedEdge(n1, n2, 1))
	if g1.Equal(g2) || g2.Equal(g1) {
		t.Errorf("graphs should not equal")
	}
}
