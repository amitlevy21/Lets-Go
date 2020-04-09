package network

import (
	"fmt"

	"gonum.org/v1/gonum/graph"
)

// Network is a graph of stations
type Network struct {
	g *Graph
}

// NewNetwork creates a network
func NewNetwork() Network {
	return Network{NewGraph()}
}

// Station returns the station from the network
// returns error if the station doesn't exist
func (n *Network) Station(id int64) (Station, error) {
	v := n.g.Node(id)
	if v == nil {
		return Station{}, fmt.Errorf("station %d not found in network", id)
	}
	return Station{id: v.ID()}, nil
}

// AddStation adds a station to the network, returns error if the station already exist
func (n *Network) AddStation(s *Station) error {
	if exist := n.g.Node(s.id); exist != nil {
		return fmt.Errorf("station %#v already exist", exist)
	}
	node := n.g.NewNode()
	s.id = node.ID()
	n.g.AddNode(node)

	return nil
}

// CheckReachability whether dst is reachable from src
func (n *Network) CheckReachability(src *Station, dst *Station) (bool, error) {
	if _, err := n.Station(src.id); err != nil {
		errMsg := "cannot check reachability. Source station %v not in network"
		return false, fmt.Errorf(errMsg, src)
	}
	if _, err := n.Station(dst.id); err != nil {
		return false, nil
	}
	found := n.g.BFS(src.id, func(node graph.Node, depth int) bool {
		return node.ID() == dst.id
	})

	return found != nil, nil
}

// ConnectStations marks that dst staion is reachable from src station.
// returns error if src or dst are nil.
func (n *Network) ConnectStations(src, dst *Station, duration float64) error {
	nSrc := n.g.Node(src.id)
	nDst := n.g.Node(dst.id)
	if nSrc == nil || nDst == nil {
		return fmt.Errorf("cannot connect stations %v to %v", nSrc, nDst)
	}
	if nSrc.ID() == nDst.ID() {
		return fmt.Errorf("cannot connect station to itself. ID: %d", nSrc.ID())
	}
	if duration < 0 {
		return fmt.Errorf("duration must be non negative. got: %f", duration)
	}
	n.g.SetWeightedEdge(n.g.NewWeightedEdge(nSrc, nDst, duration))
	return nil
}

// ValidateRoute checks if the given route is visitable in the given order.
func (n *Network) ValidateRoute(route []int64) bool {
	if len(route) == 0 {
		return false
	}
	if !n.allStationExist(route) {
		return false
	}
	for i, r := range route[1:] {
		reach, err := n.CheckReachability(&Station{id: route[i]}, &Station{id: r})
		if err != nil || !reach {
			return false
		}
	}
	return true
}

func (n *Network) allStationExist(ids []int64) bool {
	for _, id := range ids {
		_, err := n.Station(id)
		if err != nil {
			return false
		}
	}
	return true
}
