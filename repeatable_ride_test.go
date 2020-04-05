package main

import "testing"

func TestCreateRepeatableRide(t *testing.T) {
	_, err := NewRepeatable([]uint64{1, 2, 3})
	if err != nil {
		t.Errorf("error create repeatable: %s", err)
	}
}

func TestCreateWithNilStations(t *testing.T) {
	rr, err := NewRepeatable(nil)
	if err != nil {
		t.Errorf("must be able to create with nil stations")
	}
	if rr.stationsIds == nil {
		t.Error("stationIds not allocated automatically")
	}
}

func TestCreateWithConsecutiveStations(t *testing.T) {
	if _, err := NewRepeatable([]uint64{1, 2, 2}); err == nil {
		t.Errorf("cannot create with equally consecutive stations")
	}
}

func TestEqualConsecutiveStaions(t *testing.T) {
	var stationID = uint64(3)
	rr := repeatableRide{stationsIds: []uint64{1, 2, 3}}
	err := rr.AddStation(stationID)
	if err == nil {
		t.Error("adding same station twice in a row is forbidden")
	}
}

func TestAddStationToRide(t *testing.T) {
	var stationID = uint64(2)
	rr := repeatableRide{stationsIds: []uint64{1, 2, 3}}
	err := rr.AddStation(stationID)
	if err != nil {
		t.Errorf("error while adding station to ride: %s", err)
	}
	if rr.stationsIds[len(rr.stationsIds)-1] != stationID {
		t.Errorf("station id %d was not added to ride", stationID)
	}
}
