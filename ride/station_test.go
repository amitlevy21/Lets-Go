// Copyright (c) 2020 Amit Levy
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package ride

import "testing"

func TestCreateStation(t *testing.T) {
	s := Station{
		location: coordinates{
			x: 16.5,
			y: 17.5,
		},
		rideIds: map[int64]bool{1: true, 2: true, 3: true},
	}
	t.Logf("station created! %v", s)
}

func TestAddExistRide(t *testing.T) {
	rideID := int64(1)
	s := Station{
		location: coordinates{
			x: 16.5,
			y: 17.5,
		},
		rideIds: map[int64]bool{rideID: true},
	}

	if err := s.AddRide(rideID); err == nil {
		t.Error("should not add already exist")
	}
}

func TestAddRide(t *testing.T) {
	rideID := int64(5)
	s := Station{
		location: coordinates{
			x: 16.5,
			y: 17.5,
		},
		rideIds: map[int64]bool{1: true},
	}

	if err := s.AddRide(rideID); err != nil {
		t.Errorf("error while adding ride %d, err: %s", rideID, err)
	}
	if !s.rideIds[rideID] {
		t.Errorf("ride with id %d not added. %v", rideID, s.rideIds)
	}
}
