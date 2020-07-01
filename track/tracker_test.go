// Copyright (c) 2020 Amit Levy
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package track

import (
	"testing"

	"github.com/amitlevy21/lets-go/ride"
)

func TestCreate(t *testing.T) {
	tr := Tracker{}
	t.Logf("tracker %v created", tr)
}

func TestHasArrived(t *testing.T) {
	st := &ride.Station{}
	r := &ride.Ride{Route: ride.Route{StationsIds: []int64{st.ID}}}
	tr := Tracker{}
	arrived, err := tr.HasArrived(r, st)
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Log(arrived)
}

func TestFutureArrivalTime(t *testing.T) {
	st := &ride.Station{}
	r := &ride.Ride{Route: ride.Route{StationsIds: []int64{st.ID}}}
	tr := Tracker{}
	time := tr.FutureArrivalTime(r, st)
	t.Log(time)
}

func TestPastArrivalTime(t *testing.T) {
	st := &ride.Station{}
	r := &ride.Ride{Route: ride.Route{StationsIds: []int64{st.ID}}}
	tr := Tracker{}
	time := tr.PastArrivalTime(r, st)
	t.Log(time)
}

func TestNotifyWhenArrive(t *testing.T) {
	r := &ride.Ride{}
	s := &ride.Station{}
	tr := Tracker{}
	_, err := tr.NotifyWhenArrive(r, s)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestRideCapacity(t *testing.T) {
	r :=&ride.Ride{}
	tr := Tracker{}
	_, err := tr.RideCapacity(r)
	if err != nil {
		t.Errorf(err.Error())
	}
}
