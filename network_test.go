package main

import "testing"

func TestCreateNetwork(t *testing.T) {
	n := NewNetwork()
	t.Logf("network created! %v", n)
}

func TestGetNonExistStation(t *testing.T) {
	n := NewNetwork()
	_, err := n.Station(1)
	if err == nil {
		t.Error("cannot get a station that doesn't exist in the network")
	}
}

func TestAddDuplicateStation(t *testing.T) {
	n := NewNetwork()
	s := &station{id: 1}
	failIfError(t, n.AddStation(s))
	err := n.AddStation(s)
	if err == nil {
		t.Errorf("cannot add duplicate station to same network")
	}
}

func TestAddStation(t *testing.T) {
	n := NewNetwork()
	err := n.AddStation(&station{id: 0})
	if err != nil {
		t.Errorf("error while adding station to network: %s", err)
	}
	s, err := n.Station(0)
	if err != nil {
		t.Errorf("failed to get added station %v to network", s)
	}
}

func TestAddMultipleStations(t *testing.T) {
	n := NewNetwork()
	err := n.AddStation(&station{id: 0})
	if err != nil {
		t.Error("failed to add first station to network")
	}
	err = n.AddStation(&station{id: 1})
	if err != nil {
		t.Error("failed to add second station to network")
	}
	s, err := n.Station(0)
	if err != nil {
		t.Error("failed to get first station")
	}
	if s.id != 0 {
		t.Errorf("expected id %d got %d", 0, s.id)
	}
	s, err = n.Station(1)
	if err != nil {
		t.Error("failed to get second station")
	}
	if s.id != 1 {
		t.Errorf("expected id %d got %d", 1, s.id)
	}
}

func TestReachFromNonExistSourceStation(t *testing.T) {
	n := NewNetwork()
	src := &station{id: 1}
	dst := &station{id: 2}

	failIfError(t, n.AddStation(dst))
	reach, err := n.CheckReachability(src, dst)
	if reach || err == nil {
		t.Errorf("should not reach a source station not in network")
	}
}

func TestReachFromNonExistDestinationStation(t *testing.T) {
	n := NewNetwork()
	src := &station{id: 1}
	dst := &station{id: 2}

	failIfError(t, n.AddStation(src))
	reach, err := n.CheckReachability(src, dst)
	failIfError(t, err)
	if reach {
		t.Errorf("expected to not reach")
	}
}

func TestNotReach(t *testing.T) {
	n := NewNetwork()
	src := &station{id: 1}
	dst := &station{id: 2}

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
	src := &station{id: 1}
	dst := &station{id: 2}

	failIfError(t, n.AddStation(src))
	failIfError(t, n.AddStation(dst))
	failIfError(t, n.ConnectStations(src, dst))

	reach, err := n.CheckReachability(src, dst)
	failIfError(t, err)
	if !reach {
		t.Errorf("expected to reach")
	}
}

func TestReachIndirectly(t *testing.T) {
	n := NewNetwork()
	src := &station{id: 1}
	mid := &station{id: 2}
	dst := &station{id: 3}

	failIfError(t, n.AddStation(src))
	failIfError(t, n.AddStation(mid))
	failIfError(t, n.AddStation(dst))
	failIfError(t, n.ConnectStations(src, mid))
	failIfError(t, n.ConnectStations(mid, dst))

	reach, err := n.CheckReachability(src, dst)
	failIfError(t, err)
	if !reach {
		t.Errorf("expected to reach")
	}
}

func TestConnectBetweenStationsNotExist(t *testing.T) {
	n := NewNetwork()
	src := &station{id: 1}
	dst := &station{id: 2}
	err := n.ConnectStations(src, dst)
	if err == nil {
		t.Error("expected not to connect between station that don't exist")
	}
}

func TestConnectStationToItself(t *testing.T) {
	n := NewNetwork()
	src := &station{id: 1}
	failIfError(t, n.AddStation(src))
	err := n.ConnectStations(src, src)
	if err == nil {
		t.Error("expected not to connect to itself")
	}
}

func TestConnectBetweenStations(t *testing.T) {
	n := NewNetwork()
	src := &station{id: 1}
	dst := &station{id: 2}
	failIfError(t, n.AddStation(src))
	failIfError(t, n.AddStation(dst))
	failIfError(t, n.ConnectStations(src, dst))
	reach, err := n.CheckReachability(src, dst)
	failIfError(t, err)
	if !reach {
		t.Error("expected to reach after connected")
	}
}

func TestValidateEmptyRoute(t *testing.T) {
	n := NewNetwork()
	route := []int64{}
	valid := n.ValidateRoute(route)
	if valid {
		t.Errorf("should not be valid")
	}
}

func TestValidateRouteStartNotExist(t *testing.T) {
	n := NewNetwork()
	route := []int64{1}
	valid := n.ValidateRoute(route)
	if valid {
		t.Errorf("should return error stating the route is invalid")
	}
}


func TestValidateUnreachable(t *testing.T) {
	n := NewNetwork()
	s1 := &station{id: 1}
	s2 := &station{id: 2}
	failIfError(t, n.AddStation(s1))
	failIfError(t, n.AddStation(s2))
	route := []int64{s1.id, s2.id}
	valid := n.ValidateRoute(route)
	if valid {
		t.Errorf("expected route to be valid")
	}
}

func TestValidateRoute(t *testing.T) {
	n := NewNetwork()
	s1 := &station{id: 1}
	s2 := &station{id: 2}
	s3 := &station{id: 3}
	s4 := &station{id: 4}
	failIfError(t, n.AddStation(s1))
	failIfError(t, n.AddStation(s2))
	failIfError(t, n.AddStation(s3))
	failIfError(t, n.AddStation(s4))
	failIfError(t, n.ConnectStations(s1, s2))
	failIfError(t, n.ConnectStations(s2, s3))
	failIfError(t, n.ConnectStations(s2, s4))
	route := []int64{s1.id, s2.id, s3.id}
	valid := n.ValidateRoute(route)
	if !valid {
		t.Errorf("expected route to be valid")
	}
}

func TestValidateDirectRoute(t *testing.T) {
	n := NewNetwork()
	s1 := &station{id: 1}
	s2 := &station{id: 2}
	s3 := &station{id: 3}
	s4 := &station{id: 4}
	failIfError(t, n.AddStation(s1))
	failIfError(t, n.AddStation(s2))
	failIfError(t, n.AddStation(s3))
	failIfError(t, n.AddStation(s4))
	failIfError(t, n.ConnectStations(s1, s2))
	failIfError(t, n.ConnectStations(s2, s3))
	failIfError(t, n.ConnectStations(s2, s4))
	route := []int64{s1.id, s2.id, s4.id}
	valid := n.ValidateRoute(route)
	if !valid {
		t.Errorf("expected route to be valid")
	}
}

func failIfError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("error: %s", err)
	}
}
