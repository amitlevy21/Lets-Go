package network

import "testing"

func TestCreateStation(t *testing.T) {
	s := Station{
		location: coordinates{
			x: 16.5,
			y: 17.5,
		},
		rideIds: map[int64]bool{1: true, 2: true, 3: true},
	}

	t.Logf("Station created! %v", s)
}

func TestAddExistRide(t *testing.T) {
	var rideID = int64(1)
	s := Station{
		location: coordinates{
			x: 16.5,
			y: 17.5,
		},
		rideIds: map[int64]bool{rideID: true},
	}

	err := s.AddRide(rideID)
	t.Logf("AddRide : %s", err)
	if err == nil {
		t.Error("Adding a duplicate rideID should result in error")
	}
}

func TestAddRide(t *testing.T) {
	s := Station{
		location: coordinates{
			x: 16.5,
			y: 17.5,
		},
		rideIds: map[int64]bool{1: true},
	}
	var rideID = int64(5)
	err := s.AddRide(rideID)
	if err != nil {
		t.Errorf("Error while adding ride %d, err: %s", rideID, err)
	}

	if s.rideIds[rideID] != true {
		t.Errorf("Ride with id %d not added. %v", rideID, s.rideIds)
	}
}
