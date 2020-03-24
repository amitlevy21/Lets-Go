package main

import "time"

type cycle int8

const (
	hourly cycle = iota
	weekly
	yearly
)

type schedule struct {
	id                  uint64
	source              station
	destination         station
	duration            time.Time
	leaveFromSource     time.Time
	arriveAtDestination time.Time
	repeats             cycle
}
