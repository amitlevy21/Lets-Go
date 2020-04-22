// Copyright (c) 2020 Amit Levy
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package ride

import (
	"github.com/amitlevy21/lets-go/network"
	"testing"
)

func TestCreateRoute(t *testing.T) {
	r := Route{StationsIds: []int64{1, 2, 3}}
	t.Logf("route created %#v", r)
}

func TestEqualConsecutiveStaions(t *testing.T) {
	stationID := int64(3)
	n := network.NewNetwork()
	r := Route{StationsIds: []int64{1, 2, 3}}
	if err := r.AddStation(stationID, &n); err == nil {
		t.Error("adding same station twice in a row is forbidden")
	}
}

func TestAddStation(t *testing.T) {
	stationID := int64(2)
	n := network.NewNetwork()
	r := Route{StationsIds: []int64{1, 2, 3}}
	if err := r.AddStation(stationID, &n); err != nil {
		t.Errorf("error while adding station: %s", err)
	}
	if r.StationsIds[len(r.StationsIds)-1] != stationID {
		t.Errorf("station id %d was not added", stationID)
	}
}
