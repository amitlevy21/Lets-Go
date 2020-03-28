package main

import (
	"fmt"

	"github.com/starwander/goraph"
)

// Graph is a type alias for modularity
// incase we need to change the underline package that
// manages graphs, we only need to change it here
type Graph = goraph.Graph

type myVertex struct {
	id     uint64
	outTo  map[uint64]float64
	inFrom map[uint64]float64
}

type myEdge struct {
	from   uint64
	to     uint64
	weight float64
}

// NewGraph must be used to initialize a graph
func NewGraph(vertices []*myVertex) *Graph {
	g := goraph.NewGraph()
	for _, v := range vertices {
		g.AddVertexWithEdges(v)
	}
	return g
}

func (vertex *myVertex) ID() goraph.ID {
	return vertex.id
}

func (vertex *myVertex) Edges() []goraph.Edge {
	edges := make([]goraph.Edge, len(vertex.outTo)+len(vertex.inFrom))
	i := 0
	for to, weight := range vertex.outTo {
		edges[i] = &myEdge{vertex.id, to, weight}
		i++
	}
	for from, weight := range vertex.inFrom {
		edges[i] = &myEdge{from, vertex.id, weight}
		i++
	}
	return edges
}

func (edge *myEdge) Get() (goraph.ID, goraph.ID, float64) {
	return edge.from, edge.to, edge.weight
}

// BFS breadth-first-search algorithm, starts from given start node
// invokes fn function when visiting every node, if fn return true
// the function stops and returns true
func BFS(g *Graph, startVertexID uint64, fn func(v *myVertex) bool) (bool, error) {
	start, err := g.GetVertex(startVertexID)
	if err != nil {
		msg := "Cannot BFS from vertex %v. It's not in the graph"
		return false, fmt.Errorf(msg, startVertexID)
	}
	if err := g.CheckIntegrity(); err != nil {
		return false, fmt.Errorf("Graph %v is not valid. %s", g, err)
	}
	s := start.(*myVertex)
	visited := make(map[uint64]bool)
	q := []uint64{}
	visited[s.id] = true
	q = append(q, s.id)
	if stop := fn(s); stop {
		return true, nil
	}

	for len(q) > 0 {
		vID := q[0]
		q = q[1:]
		v, _ := g.GetVertex(vID)
		vc := v.(*myVertex)
		for _, e := range vc.Edges() {
			_, to, _ := e.Get()
			toID := to.(uint64)
			if visited[toID] == false {
				q = append(q, toID)
				toV, _ := g.GetVertex(toID)
				visited[toID] = true
				if stop := fn(toV.(*myVertex)); stop {
					return true, nil
				}
			}
		}
	}

	return true, nil
}
