// Copyright (c) 2020 Amit Levy
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package ride

import (
	"fmt"
	"testing"
)

func TestCreateRide(t *testing.T) {
	r := Ride{
		rStatus:        pending,
		numPassengers:  0,
		availableSeats: 30,
		crewMembersIds: []int64{1, 2, 3},
		vehicle:        int64(0),
		latestStation:  int64(0),
	}
	t.Logf("ride created! %v", r)
}

type testCase struct {
	name     string
	or       *Ride
	fn       string
	expected status
	fail     bool
}

var testCases = []testCase{
	{
		"newly with pending status",
		&Ride{},
		"",
		pending,
		false,
	},
	{
		"start already started",
		&Ride{rStatus: ongoing},
		"start",
		ongoing,
		true,
	},
	{
		"start finished ride",
		&Ride{rStatus: finished},
		"start",
		finished,
		true,
	},
	{
		"start cancelled ride",
		&Ride{rStatus: cancelled},
		"start",
		cancelled,
		true,
	},
	{
		"start pending ride",
		&Ride{},
		"start",
		ongoing,
		false,
	},
	{
		"finish finished ride",
		&Ride{rStatus: finished},
		"finish",
		finished,
		true,
	},
	{
		"finish pending ride",
		&Ride{},
		"finish",
		finished,
		false,
	},
	{
		"finish cancelled ride",
		&Ride{rStatus: cancelled},
		"finish",
		cancelled,
		true,
	},
	{
		"finish ongoing ride",
		&Ride{rStatus: ongoing},
		"finish",
		finished,
		false,
	},
	{
		"cancel pending ride",
		&Ride{rStatus: pending},
		"cancel",
		cancelled,
		false,
	},
	{
		"cancel finished ride",
		&Ride{rStatus: finished},
		"cancel",
		cancelled,
		false,
	},
	{
		"cancel ongoing ride",
		&Ride{rStatus: ongoing},
		"cancel",
		cancelled,
		false,
	},
	{
		"cancel cancelled ride",
		&Ride{rStatus: cancelled},
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

func fnFromStruct(r *Ride, fn string) error {
	switch fn {
	case "start":
		return r.start()
	case "finish":
		return r.finish()
	case "cancel":
		return r.cancel()
	case "":
		return nil
	default:
		return fmt.Errorf("%T: no method found called %s", r, fn)
	}
}

func TestStatusString(t *testing.T) {
	s := fmt.Sprint(pending)
	expected := "pending"
	if s != expected {
		t.Errorf("expected %s to be %s", s, expected)
	}
}
