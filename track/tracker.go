// Copyright (c) 2020 Amit Levy
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package track

import (
	"time"

	"github.com/amitlevy21/lets-go/ride"
)

// Tracker manages the state of ride execution.
type Tracker struct {}

// HasArrived returns whether the ride has reached the station.
// returns error if ride's route those not include the station
func (t *Tracker) HasArrived(r *ride.Ride, s *ride.Station) (bool, error) {
	return false, nil
}

// FutureArrivalTime returns the when the ride will reach the station.
func (t *Tracker) FutureArrivalTime(r *ride.Ride, s *ride.Station) time.Time {
	return time.Now()
}

// PastArrivalTime returns the when the ride had arrived the station.
func (t *Tracker) PastArrivalTime(r *ride.Ride, s *ride.Station) time.Time {
	return time.Now()
}

// NotifyWhenArrive send to a channel when ride will reach station.
func (t *Tracker) NotifyWhenArrive(r *ride.Ride, s *ride.Station) (chan bool, error) {
	return make(chan bool), nil
}

// RideCapacity returns how many people are in the current ride.
// returns error if the ride was not found or has not started.
func (t *Tracker) RideCapacity(r *ride.Ride) (uint32, error) {
	return 0, nil
}
