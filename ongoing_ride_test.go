package main

import "testing"

func TestCreateOngoingRide(t *testing.T) {
	or := ongoingRide{
		rStatus:        pending,
		numPassengers:  0,
		availableSeats: 30,
		crewMembersIds: []uint64{1, 2, 3},
		rVehicle:       vehicle{},
		latestStation:  station{},
	}

	t.Logf("Ongoing ride created! %v", or)
}
