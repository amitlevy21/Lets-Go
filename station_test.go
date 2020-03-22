package main

import "testing"

func TestCreateStation(t *testing.T) {
	s := station{
		location: coordinates{
			x: 16.5,
			y: 17.5,
		},
		rideIds: []int64{1, 2, 3},
	}

	t.Logf("Station created! %v", s)
}
