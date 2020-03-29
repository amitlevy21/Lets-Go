package main

import "fmt"

// Network is a graph of stations
type Network struct {
	g *Graph
}

// NewNetwork creates a network
func NewNetwork() Network {
	return Network{NewGraph([]*myVertex{})}
}

// GetStation returns the station from the network
// returns error if the station doesn't exist
func (n *Network) GetStation(id uint64) (station, error) {
	v, err := n.g.GetVertex(id)
	if v == nil {
		return station{}, fmt.Errorf("station %d not found in network", id)
	}
	return station{id: v.(*myVertex).id}, err
}

// AddStation adds a vertex to the network
func (n *Network) AddStation(s *station) error {
	return n.g.AddVertex(s.id, &myVertex{id: s.id})
}

// CheckReachability whether dst is reachable from src
func (n *Network) CheckReachability(src *station, dst *station) (bool, error) {
	if _, err := n.GetStation(src.id); err != nil {
		errMsg := "Cannot check reachability. Source station %v not in network"
		return false, fmt.Errorf(errMsg, src)
	}
	if _, err := n.GetStation(dst.id); err != nil {
		return false, nil
	}
	found := false
	_, err := BFS(n.g, src.id, func(v *myVertex) bool {
		if v.id == dst.id {
			found = true
			return true
		}
		return false
	})
	if err != nil {
		return false, err
	}

	return found, nil
}

// ConnectStations marks that dst staion is reachable from src station
func (n *Network) ConnectStations(src *station, dst *station) error {
	if err := n.hackSyncVertexData(src, dst); err != nil {
		return err
	}
	return n.g.AddEdge(src.id, dst.id, 0, nil)
}

// Ugly hack need because there is no exported method from the graph
// package so we can access the edges of a vertex.
// The current method v.Edges() is not synced with internal struct data of the packaged
func (n *Network) hackSyncVertexData(src *station, dst *station) error {
	sv, err := n.g.GetVertex(src.id)
	dv, err := n.g.GetVertex(dst.id)
	if err != nil {
		return fmt.Errorf("Cannot connect between station not in network")
	}
	svv := sv.(*myVertex)
	dvv := dv.(*myVertex)
	if svv.outTo == nil {
		svv.outTo = make(map[uint64]float64)
	}
	if dvv.inFrom == nil {
		dvv.inFrom = make(map[uint64]float64)
	}
	svv.outTo[dvv.id] = 0
	dvv.inFrom[svv.id] = 0
	return nil
}
