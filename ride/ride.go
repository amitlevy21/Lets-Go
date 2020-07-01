// Copyright (c) 2020 Amit Levy
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package ride

import "fmt"

type status int8

const (
	pending status = iota
	ongoing
	finished
	cancelled
)

var statusStr = []string{"pending", "ongoing", "finished", "cancelled"}

// Ride is a moving vehicle with passengers.
type Ride struct {
	id              int64
	rStatus         status
	numPassengers   int32
	availableSeats  int32
	crewMembersIds  []int64
	vehicle         int64
	latestStation   int64
	Route           Route
	VisitedStations []int64
}

func (r *Ride) start() error {
	if r.rStatus != pending {
		return fmt.Errorf("ride already started")
	}
	r.rStatus = ongoing
	return nil
}

func (r *Ride) finish() error {
	if r.rStatus == finished {
		return fmt.Errorf("ride already finished")
	}
	if r.rStatus == cancelled {
		return fmt.Errorf("cannot finish cancelled ride")
	}
	r.rStatus = finished
	return nil
}

func (r *Ride) cancel() error {
	r.rStatus = cancelled
	return nil
}

func (s status) String() string {
	return statusStr[s]
}
