package main

type status int8

const (
	pending status = iota
	ongoing
	finished
)

type ongoingRide struct {
	id             int64
	rStatus        status
	numPassengers  int32
	availableSeats int32
	crewMembersIds []int64
	rVehicle       vehicle
	latestStation  station
}
