package main

import "fmt"

type coordinates struct {
	x float64
	y float64
}

type station struct {
	id       uint64
	location coordinates
	rideIds  map[uint64]bool
}

// AddRide adds a ride to the station, returns error if rideId exists
func (s *station) AddRide(rideID uint64) error {
	if s.rideIds[rideID] == true {
		return fmt.Errorf("Cannot add rideID %d. Already exists", rideID)
	}
	s.rideIds[rideID] = true
	return nil
}
