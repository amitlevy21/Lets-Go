package main

import "testing"

type neighbors = map[uint64]float64

func TestCreateGraph(t *testing.T) {
	g := NewGraph([]*myVertex{&myVertex{id: 1}, &myVertex{id: 2}})
	t.Logf("Graph %v created!", g)
}

func TestBFSFromNonExistSrc(t *testing.T) {
	g := NewGraph([]*myVertex{})
	if _, err := BFS(g, 1, func(v *myVertex) bool {return true}); err == nil {
		t.Error("Should not BFS when src node not in graph")
	}
}

var BFSTests = []struct {
	name          string
	g             *Graph
	startVertexID uint64
	sum           uint64
}{
	{
		"Single vertex",
		NewGraph([]*myVertex{&myVertex{id: 1}}),
		1,
		1,
	},
	{
		"Single vertex with self pointing edge",
		NewGraph([]*myVertex{&myVertex{id: 1, outTo: neighbors{1: 1.0}, inFrom: neighbors{1: 1.0}}}),
		1,
		1,
	},
	{
		"Two unconnected vertices",
		NewGraph([]*myVertex{&myVertex{id: 1}, &myVertex{id: 2}}),
		1,
		1,
	},
	{
		"Two connected vertices",
		NewGraph([]*myVertex{&myVertex{id: 1, outTo: neighbors{2: 1.0}}, &myVertex{id: 2}}),
		1,
		3,
	},
}

func TestBFS(t *testing.T) {
	for _, tt := range BFSTests {
		t2 := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			sum := uint64(0)
			_, err := BFS(t2.g, t2.startVertexID, func(v *myVertex) bool {
				t.Logf("Adding %d to sum %d", v.id, sum)
				sum += v.id
				return false
			})
			if err != nil {
				t.Errorf("BFS failed: %s", err)
			}
			if sum != t2.sum {
				t.Errorf("Expected sum of visits %d, but got %d", t2.sum, sum)
			}
		})
	}
}

func TestBFSEarlyStop(t *testing.T) {
	g := NewGraph([]*myVertex{&myVertex{id: 1}, &myVertex{id: 2}})
	sum := uint64(0)
	early, err := BFS(g, 1, func(v *myVertex) bool {
		sum += v.id
		return true
	})
	if err != nil {
		t.Error(err)
	}
	if !early {
		t.Error("Expected to stop early")
	}
	if sum != 1 {
		t.Error("Did not iterate before early stopping")
	}
}
