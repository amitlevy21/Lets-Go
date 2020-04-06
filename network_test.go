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
	n.AddStation(s)
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
	_, err = n.Station(0)
	if err != nil {
		t.Error("failed to get first station")
	}
	_, err = n.Station(1)
	if err != nil {
		t.Error("failed to get second station")
	}
}

func TestReachFromNonExistSourceStation(t *testing.T) {
	n := NewNetwork()
	src := &station{id: 1}
	dst := &station{id: 2}

	n.AddStation(dst)
	reach, err := n.CheckReachability(src, dst)
	if reach || err == nil {
		t.Errorf("should not be able to reach a source station that was not added to network")
	}
}

func TestReachFromNonExistDestinationStation(t *testing.T) {
	n := NewNetwork()
	src := &station{id: 1}
	dst := &station{id: 2}

	n.AddStation(src)
	reach, err := n.CheckReachability(src, dst)
	if err != nil {
		t.Errorf("expected to checkReachability without errors %s", err)
	}
	if reach {
		t.Errorf("expected to not reach without errors")
	}
}

func TestNotReach(t *testing.T) {
	n := NewNetwork()
	src := &station{id: 1}
	dst := &station{id: 2}

	n.AddStation(src)
	n.AddStation(dst)
	reach, err := n.CheckReachability(src, dst)
	if err != nil {
		t.Errorf("expected to checkReachability without errors %s", err)
	}
	if reach {
		t.Errorf("expected to not reach")
	}
}

func TestReachDirectly(t *testing.T) {
	n := NewNetwork()
	src := &station{id: 1}
	dst := &station{id: 2}

	n.AddStation(src)
	n.AddStation(dst)
	n.ConnectStations(src, dst)

	reach, err := n.CheckReachability(src, dst)
	if err != nil {
		t.Errorf("expected to checkReachability without errors %s", err)
	}
	if !reach {
		t.Errorf("expected to reach")
	}
}

func TestReachIndirectly(t *testing.T) {
	n := NewNetwork()
	src := &station{id: 1}
	mid := &station{id: 2}
	dst := &station{id: 3}

	n.AddStation(src)
	n.AddStation(mid)
	n.AddStation(dst)
	n.ConnectStations(src, mid)
	n.ConnectStations(mid, dst)

	reach, err := n.CheckReachability(src, dst)
	if err != nil {
		t.Errorf("unexpected error %s", err)
	}
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
func TestConnectBetweenStations(t *testing.T) {
	n := NewNetwork()
	src := &station{id: 1}
	dst := &station{id: 2}
	n.AddStation(src)
	n.AddStation(dst)
	err := n.ConnectStations(src, dst)
	if err != nil {
		t.Error("expect to connect stations")
	}
	reach, err := n.CheckReachability(src, dst)
	if !reach || err != nil {
		t.Error("expected to reach after connected")
	}
}
