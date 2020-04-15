// Copyright (c) 2020 Amit Levy
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package network

import (
	"fmt"

	"gonum.org/v1/gonum/graph"
)

// Network is a graph of elements
type Network struct {
	g *Graph
	m map[int64]int64
}

// NewNetwork creates a network
func NewNetwork() Network {
	return Network{NewGraph(), make(map[int64]int64)}
}

// Get returns the corresponding nodeID created when elementID was added.
func (n *Network) Get(elementID int64) (nodeID int64, ok bool) {
	nodeID, ok = n.m[elementID]
	return nodeID, ok
}

// Add adds an element to the network, returns error if element already exist.
func (n *Network) Add(id int64) error {
	if _, ok := n.Get(id); ok {
		return fmt.Errorf("element %d already exist", id)
	}
	node := n.g.NewNode()
	n.g.AddNode(node)
	n.m[id] = node.ID()

	return nil
}

// CheckReachability checks if you can reach dst from src.
func (n *Network) CheckReachability(src, dst int64) (bool, error) {
	if _, ok := n.Get(src); !ok {
		return false, fmt.Errorf("reachability: src %d not exist", src)
	}
	if _, ok := n.Get(dst); !ok {
		return false, nil
	}
	found := n.g.BFS(n.m[src], func(node graph.Node, depth int) bool {
		return node.ID() == n.m[dst]
	})

	return found != nil, nil
}

// Connect marks dst as reachable from src.
// returns error if src or dst are nil.
func (n *Network) Connect(srcID, dstID int64, duration float64) error {
	src := n.g.Node(n.m[srcID])
	dst := n.g.Node(n.m[dstID])
	if src == nil || dst == nil {
		return fmt.Errorf("cannot connect %v to %v", src, dst)
	}
	if src.ID() == dst.ID() {
		return fmt.Errorf("cannot connect element to itself. ID: %d", src.ID())
	}
	if duration < 0 {
		return fmt.Errorf("duration must be non negative. got: %f", duration)
	}
	n.g.SetWeightedEdge(n.g.NewWeightedEdge(src, dst, duration))
	return nil
}

// ValidatePath checks if the given path p is visitable in the given order.
func (n *Network) ValidatePath(p []int64) (valid bool, reason error) {
	if len(p) == 0 {
		return false, fmt.Errorf("route cannot be zero length")
	}
	if ok, guilty := n.allExist(p); !ok {
		return false, fmt.Errorf("%d not in network", guilty)
	}
	for i, s := range p[1:] {
		if reach, err := n.CheckReachability(p[i], s); err != nil || !reach {
			return false, fmt.Errorf("cannot reach from %d to %d", p[i], s)
		}
	}
	return true, nil
}

func (n *Network) allExist(ids []int64) (ok bool, guilty int64) {
	for _, id := range ids {
		if e, ok := n.Get(id); !ok {
			return false, e
		}
	}
	return true, 0
}
