package main

import "testing"

func TestCreateRepeatableRide(t *testing.T) {
	rr := repeatableRide{
		stationsIds:  []int64{1, 2, 3},
		schedulesIds: []int64{1, 2, 3},
	}

	t.Logf("Repeatable ride created! %v", rr)
}
