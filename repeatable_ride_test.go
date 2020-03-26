package main

import "testing"

func TestCreateRepeatableRide(t *testing.T) {
	rr := repeatableRide{
		stationsIds:  []uint64{1, 2, 3},
		schedulesIds: map[uint64]bool{1: true, 2: true, 3: true},
	}

	t.Logf("Repeatable ride created! %v", rr)
}

func TestCreateRoute(t *testing.T) {
	r := NewRoute()

	t.Logf("Route created! %v", r)
}

func TestAddExistSchedule(t *testing.T) {
	var scheduleID uint64 = 1
	rr := repeatableRide{
		stationsIds:  []uint64{1, 2, 3},
		schedulesIds: map[uint64]bool{scheduleID: true, 2: true, 3: true},
	}

	err := rr.AddSchedule(scheduleID)
	if err == nil {
		t.Error("Adding a schedule with the same id is forbidden")
	}
}

func TestAddSchedule(t *testing.T) {
	var scheduleID uint64 = 2
	rr := repeatableRide{
		stationsIds:  []uint64{1, 2, 3},
		schedulesIds: map[uint64]bool{1: true},
	}
	err := rr.AddSchedule(scheduleID)
	if err != nil {
		t.Errorf("Error while adding schedule to ride: %s", err)
	}
	if rr.schedulesIds[scheduleID] != true {
		t.Errorf("Schedule id %d was not added to ride", scheduleID)
	}
}
