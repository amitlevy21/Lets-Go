package main

import "fmt"

type coordinates struct {
	x float64
	y float64
}

// TODO: Add NewStation function to unsure unique id
type station struct {
	id       int64
	location coordinates
	rideIds  map[int64]bool
}

// AddRide adds a ride to the station, returns error if rideId exists
func (s *station) AddRide(rideID int64) error {
	if s.rideIds[rideID] {
		return fmt.Errorf("Cannot add rideID %d. Already exists", rideID)
	}
	s.rideIds[rideID] = true
	return nil
}
