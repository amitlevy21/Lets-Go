package main

type status int8

const (
	pending status = iota
	ongoing
	finished
)

type ongoingRide struct {
	id             uint64
	rStatus        status
	numPassengers  int32
	availableSeats int32
	crewMembersIds []uint64
	rVehicle       vehicle
	latestStation  station
}
