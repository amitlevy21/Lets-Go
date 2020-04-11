package network

import "testing"

func TestCreateNetwork(t *testing.T) {
	n := NewNetwork()
	t.Logf("network created! %v", n)
}

func TestGetNonExistStation(t *testing.T) {
	n := NewNetwork()
	if _, ok := n.Get(1); ok {
		t.Error("cannot get a station that doesn't exist in the network")
	}
}

func TestAddDuplicateStation(t *testing.T) {
	n := NewNetwork()
	s := int64(1)
	failIfError(t, n.AddStation(s))
	if err := n.AddStation(s); err == nil {
		t.Errorf("cannot add duplicate station to same network")
	}
}

func TestAddStation(t *testing.T) {
	n := NewNetwork()
	s := int64(0)
	if err := n.AddStation(s); err != nil {
		t.Errorf("error while adding station to network: %s", err)
	}
	if _, ok := n.Get(s); !ok {
		t.Errorf("failed to get added station %d to network", s)
	}
}

func TestAddMultipleStations(t *testing.T) {
	n := NewNetwork()
	s1 := int64(0)
	s2 := int64(1)
	if err := n.AddStation(s1); err != nil {
		t.Error("failed to add first station to network")
	}
	if err := n.AddStation(s2); err != nil {
		t.Error("failed to add second station to network")
	}
	if _, ok := n.Get(s1); !ok {
		t.Error("failed to get first station")
	}
	if _, ok := n.Get(s2); !ok {
		t.Error("failed to get second station")
	}
}

func TestReachFromNonExistSourceStation(t *testing.T) {
	n := NewNetwork()
	failIfError(t, n.AddStation(2))
	reach, err := n.CheckReachability(1, 2)
	if reach || err == nil {
		t.Errorf("should not reach a source station not in network")
	}
}

func TestReachFromNonExistDestinationStation(t *testing.T) {
	n := NewNetwork()
	failIfError(t, n.AddStation(1))
	reach, err := n.CheckReachability(1, 2)
	failIfError(t, err)
	if reach {
		t.Errorf("expected to not reach")
	}
}

func TestNotReach(t *testing.T) {
	n := NewNetwork()
	src := int64(1)
	dst := int64(2)
	failIfError(t, n.AddStation(src))
	failIfError(t, n.AddStation(dst))
	reach, err := n.CheckReachability(src, dst)
	failIfError(t, err)
	if reach {
		t.Errorf("expected to not reach")
	}
}

func TestReachDirectly(t *testing.T) {
	n := NewNetwork()
	src := int64(1)
	dst := int64(2)
	failIfError(t, n.AddStation(src))
	failIfError(t, n.AddStation(dst))
	failIfError(t, n.ConnectStations(src, dst, 0))

	reach, err := n.CheckReachability(src, dst)
	failIfError(t, err)
	if !reach {
		t.Errorf("expected to reach")
	}
}

func TestReachIndirectly(t *testing.T) {
	n := NewNetwork()
	src := int64(1)
	mid := int64(2)
	dst := int64(3)

	failIfError(t, n.AddStation(src))
	failIfError(t, n.AddStation(mid))
	failIfError(t, n.AddStation(dst))
	failIfError(t, n.ConnectStations(src, mid, 0))
	failIfError(t, n.ConnectStations(mid, dst, 0))

	reach, err := n.CheckReachability(src, dst)
	failIfError(t, err)
	if !reach {
		t.Errorf("expected to reach")
	}
}

func TestConnectBetweenStationsNotExist(t *testing.T) {
	n := NewNetwork()
	src := int64(1)
	dst := int64(2)
	err := n.ConnectStations(src, dst, 0)
	if err == nil {
		t.Error("expected not to connect between station that don't exist")
	}
}

func TestConnectStationToItself(t *testing.T) {
	n := NewNetwork()
	src := int64(1)
	failIfError(t, n.AddStation(src))
	err := n.ConnectStations(src, src, 0)
	if err == nil {
		t.Error("expected not to connect to itself")
	}
}

func TestSetConnectStationNegDuration(t *testing.T) {
	n := NewNetwork()
	s1 := int64(1)
	s2 := int64(2)
	failIfError(t, n.AddStation(s1))
	failIfError(t, n.AddStation(s2))
	err := n.ConnectStations(s1, s2, -1)
	if err == nil {
		t.Error("should not accept negative duration")
	}
}

func TestConnectBetweenStations(t *testing.T) {
	n := NewNetwork()
	src := int64(1)
	dst := int64(2)
	failIfError(t, n.AddStation(src))
	failIfError(t, n.AddStation(dst))
	failIfError(t, n.ConnectStations(src, dst, 1))
	reach, err := n.CheckReachability(src, dst)
	failIfError(t, err)
	if !reach {
		t.Error("expected to reach after connected")
	}
}

func TestValidateEmptyRoute(t *testing.T) {
	n := NewNetwork()
	route := []int64{}
	if valid, _ := n.ValidateRoute(route); valid {
		t.Errorf("should not be valid")
	}
}

func TestValidateRouteStartNotExist(t *testing.T) {
	n := NewNetwork()
	route := []int64{1}
	if valid, _ := n.ValidateRoute(route); valid {
		t.Errorf("should return error stating the route is invalid")
	}
}

func TestValidateUnreachable(t *testing.T) {
	n := NewNetwork()
	s1 := int64(1)
	s2 := int64(2)
	failIfError(t, n.AddStation(s1))
	failIfError(t, n.AddStation(s2))
	route := []int64{s1, s2}
	if valid, _ := n.ValidateRoute(route); valid {
		t.Errorf("expected route to be valid")
	}
}

func TestValidateRoute(t *testing.T) {
	n := NewNetwork()
	s1 := int64(1)
	s2 := int64(2)
	s3 := int64(3)
	s4 := int64(4)
	failIfError(t, n.AddStation(s1))
	failIfError(t, n.AddStation(s2))
	failIfError(t, n.AddStation(s3))
	failIfError(t, n.AddStation(s4))
	failIfError(t, n.ConnectStations(s1, s2, 0))
	failIfError(t, n.ConnectStations(s2, s3, 0))
	failIfError(t, n.ConnectStations(s2, s4, 0))
	route := []int64{s1, s2, s3}
	if valid, _ := n.ValidateRoute(route); !valid {
		t.Errorf("expected route to be valid")
	}
}

func TestValidateDirectRoute(t *testing.T) {
	n := NewNetwork()
	s1 := int64(1)
	s2 := int64(2)
	s3 := int64(3)
	s4 := int64(4)
	failIfError(t, n.AddStation(s1))
	failIfError(t, n.AddStation(s2))
	failIfError(t, n.AddStation(s3))
	failIfError(t, n.AddStation(s4))
	failIfError(t, n.ConnectStations(s1, s2, 0))
	failIfError(t, n.ConnectStations(s2, s3, 0))
	failIfError(t, n.ConnectStations(s2, s4, 0))
	route := []int64{s1, s2, s4}
	if valid, _ := n.ValidateRoute(route); !valid {
		t.Errorf("expected route to be valid")
	}
}

func failIfError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("unexpected error during test: %s", err)
	}
}
