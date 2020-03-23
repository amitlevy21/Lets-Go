package main

import "time"

type cycle int8

const (
	hourly cycle = iota
	weekly
	yearly
)

type schedule struct {
	source              station
	destination         station
	duration            time.Time
	leaveFromSource     time.Time
	arriveAtDestination time.Time
	repeats             cycle
}
