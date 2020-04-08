package main

import "testing"

func TestCreateRepeatableRide(t *testing.T) {
	rr := repeatableRide{stationsIds: []int64{1, 2, 3}}
	t.Logf("repeatable created %#v", rr)
}

func TestEqualConsecutiveStaions(t *testing.T) {
	var stationID = int64(3)
	rr := repeatableRide{stationsIds: []int64{1, 2, 3}}
	err := rr.AddStation(stationID)
	if err == nil {
		t.Error("adding same station twice in a row is forbidden")
	}
}

func TestAddStationToRide(t *testing.T) {
	var stationID = int64(2)
	rr := repeatableRide{stationsIds: []int64{1, 2, 3}}
	err := rr.AddStation(stationID)
	if err != nil {
		t.Errorf("error while adding station to ride: %s", err)
	}
	if rr.stationsIds[len(rr.stationsIds)-1] != stationID {
		t.Errorf("station id %d was not added to ride", stationID)
	}
}
