// Copyright (c) 2020 Amit Levy
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package ride

import "fmt"

type coordinates struct {
	x float64
	y float64
}

// Station is a stop for a ride
// TODO Add NewStation function to unsure unique id
type Station struct {
	id       int64
	location coordinates
	rideIds  map[int64]bool
}

// AddRide adds a ride to the station, returns error if rideId exists
func (s *Station) AddRide(rideID int64) error {
	if s.rideIds[rideID] {
		return fmt.Errorf("Cannot add rideID %d. Already exists", rideID)
	}
	s.rideIds[rideID] = true
	return nil
}
