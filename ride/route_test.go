package ride

import (
	"testing"
	"github.com/amitlevy21/lets-go/network"
)

func TestCreateRoute(t *testing.T) {
	r := route{stationsIds: []int64{1, 2, 3}}
	t.Logf("route created %#v", r)
}

func TestEqualConsecutiveStaions(t *testing.T) {
	stationID := int64(3)
	n := network.NewNetwork()
	r := route{stationsIds: []int64{1, 2, 3}}
	err := r.AddStation(stationID, &n)
	if err == nil {
		t.Error("adding same station twice in a row is forbidden")
	}
}

func TestAddStation(t *testing.T) {
	stationID := int64(2)
	n := network.NewNetwork()
	r := route{stationsIds: []int64{1, 2, 3}}
	if err := r.AddStation(stationID, &n); err != nil {
		t.Errorf("error while adding station: %s", err)
	}
	if r.stationsIds[len(r.stationsIds)-1] != stationID {
		t.Errorf("station id %d was not added", stationID)
	}
}