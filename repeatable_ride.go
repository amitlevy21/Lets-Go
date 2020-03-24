package main

import "fmt"

type repeatableRide struct {
	id           uint64
	stationsIds  []uint64
	schedulesIds map[uint64]bool
}

func (rr *repeatableRide) AddSchedule(scheduleID uint64) error {
	if rr.schedulesIds[scheduleID] == true {
		return fmt.Errorf("Schedule with id %d already exists", scheduleID)
	}
	rr.schedulesIds[scheduleID] = true
	return nil
}
