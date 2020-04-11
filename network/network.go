package network

import (
	"fmt"

	"gonum.org/v1/gonum/graph"
)

// Network is a graph of stations
type Network struct {
	g    *Graph
	sToN map[int64]int64
}

// NewNetwork creates a network
func NewNetwork() Network {
	return Network{NewGraph(), make(map[int64]int64)}
}

// Get returns the corresponding nodeID created when stationID was added.
func (n *Network) Get(stationID int64) (nodeID int64, ok bool) {
	nodeID, ok = n.sToN[stationID]
	return nodeID, ok
}

// AddStation adds a station to the network, returns error if the station already exist
func (n *Network) AddStation(id int64) error {
	if _, ok := n.Get(id); ok {
		return fmt.Errorf("station %d already exist", id)
	}
	node := n.g.NewNode()
	n.g.AddNode(node)
	n.sToN[id] = node.ID()

	return nil
}

// CheckReachability whether dst is reachable from src
func (n *Network) CheckReachability(srcID, dstID int64) (bool, error) {
	if src, ok := n.Get(srcID); !ok {
		return false, fmt.Errorf("reachability: src %d not exist", src)
	}
	if _, ok := n.Get(dstID); !ok {
		return false, nil
	}
	found := n.g.BFS(n.sToN[srcID], func(node graph.Node, depth int) bool {
		return node.ID() == n.sToN[dstID]
	})

	return found != nil, nil
}

// ConnectStations marks that dst staion is reachable from src station.
// returns error if src or dst are nil.
func (n *Network) ConnectStations(srcID, dstID int64, duration float64) error {
	nSrc := n.g.Node(n.sToN[srcID])
	nDst := n.g.Node(n.sToN[dstID])
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
func (n *Network) ValidateRoute(route []int64) (valid bool, reason error) {
	if len(route) == 0 {
		return false, fmt.Errorf("route cannot be zero length")
	}
	if ok, guilty := n.allStationExist(route); !ok {
		return false, fmt.Errorf("station %d not in network", guilty)
	}
	for i, s := range route[1:] {
		reach, err := n.CheckReachability(route[i], s)
		if err != nil || !reach {
			return false, fmt.Errorf("cannot reach from %d to %d", route[i], s)
		}
	}
	return true, nil
}

func (n *Network) allStationExist(ids []int64) (ok bool, guilty int64) {
	for _, id := range ids {
		s, ok := n.Get(id)
		if !ok {
			return false, s
		}
	}
	return true, 0
}
