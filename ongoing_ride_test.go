package main

import (
	"fmt"
	"testing"
)

type spec struct {
	name     string
	or       *ongoingRide
	fn       string
	expected status
	fail     bool
}

func TestCreateOngoingRide(t *testing.T) {
	or := ongoingRide{
		rStatus:        pending,
		numPassengers:  0,
		availableSeats: 30,
		crewMembersIds: []uint64{1, 2, 3},
		rVehicle:       vehicle{},
		latestStation:  station{},
	}
	t.Logf("Ongoing ride created! %v", or)
}

var testCases = []spec{
	{
		"newly with pending status",
		&ongoingRide{},
		"",
		pending,
		false,
	},
	{
		"start already started",
		&ongoingRide{rStatus: ongoing},
		"start",
		ongoing,
		true,
	},
	{
		"start finished ride",
		&ongoingRide{rStatus: finished},
		"start",
		finished,
		true,
	},
	{
		"start cancelled ride",
		&ongoingRide{rStatus: cancelled},
		"start",
		cancelled,
		true,
	},
	{
		"start pending ride",
		&ongoingRide{},
		"start",
		ongoing,
		false,
	},
	{
		"finish finished ride",
		&ongoingRide{rStatus: finished},
		"finish",
		finished,
		true,
	},
	{
		"finish pending ride",
		&ongoingRide{},
		"finish",
		finished,
		false,
	},
	{
		"finish cancelled ride",
		&ongoingRide{rStatus: cancelled},
		"finish",
		cancelled,
		true,
	},
	{
		"finish ongoing ride",
		&ongoingRide{rStatus: ongoing},
		"finish",
		finished,
		false,
	},
	{
		"cancel pending ride",
		&ongoingRide{rStatus: pending},
		"cancel",
		cancelled,
		false,
	},
	{
		"cancel finished ride",
		&ongoingRide{rStatus: finished},
		"cancel",
		cancelled,
		false,
	},
	{
		"cancel ongoing ride",
		&ongoingRide{rStatus: ongoing},
		"cancel",
		cancelled,
		false,
	},
	{
		"cancel cancelled ride",
		&ongoingRide{rStatus: cancelled},
		"cancel",
		cancelled,
		false,
	},
}

func TestRideStatus(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t2 := tt // store loop var, before changed by the for loop
			t.Parallel()
			if err := fnFromStruct(t2.or, t2.fn); err != nil {
				if !t2.fail {
					t.Errorf("%s should not fail but did. error: %s", t2.name, err)
				}
			}
			if t2.expected != t2.or.rStatus {
				t.Errorf("expected status %s but got %s", t2.expected, t2.or.rStatus)
			}
		})
	}
}

func fnFromStruct(or *ongoingRide, fn string) error {
	switch fn {
	case "start":
		return or.start()
	case "finish":
		return or.finish()
	case "cancel":
		return or.cancel()
	case "":
		return nil
	default:
		return fmt.Errorf("%T: no method found called %s", or, fn)
	}
}

func TestStatusString(t *testing.T) {
	s := fmt.Sprint(pending)
	expected := "pending"
	if s != expected {
		t.Errorf("expected %s to be %s", s, expected)
	}
}
