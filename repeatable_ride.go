package main

import (
	"fmt"
	"github.com/starwander/goraph"
)

// Route is a graph where the edges are station and the vertices are durations
type Route *goraph.Graph

type repeatableRide struct {
	id           uint64
	stationsIds  []uint64
	schedulesIds map[uint64]bool
}

// NewRoute creates a route
func NewRoute() Route {
	return goraph.NewGraph()
}

func (rr *repeatableRide) AddSchedule(scheduleID uint64) error {
	if rr.schedulesIds[scheduleID] == true {
		return fmt.Errorf("Schedule with id %d already exists", scheduleID)
	}
	rr.schedulesIds[scheduleID] = true
	return nil
}
