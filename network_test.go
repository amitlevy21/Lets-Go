package main

import "testing"

func TestCreateNetwork(t *testing.T) {
	n := NewNetwork()
	t.Logf("Network created! %v", n)
}

func TestGetNonExistStation(t *testing.T) {
	n := NewNetwork()
	_, err := n.GetStation(1)
	if err == nil {
		t.Error("You cannot get a station that doesn't exist in the network")
	}
}

func TestAddDuplicateStation(t *testing.T) {
	n := NewNetwork()
	n.AddStation(&station{id: 1})
	err := n.AddStation(&station{id: 1})
	if err == nil {
		t.Errorf("Cannot add duplicate station to same network")
	}
}

func TestAddStation(t *testing.T) {
	n := NewNetwork()
	err := n.AddStation(&station{id: 1})
	if err != nil {
		t.Errorf("Error while adding station to network: %s", err)
	}
	s, err := n.GetStation(1)
	if err != nil {
		t.Errorf("Failed to get added station %v to network", s)
	}
}

func TestAddMultipleStations(t *testing.T) {
	n := NewNetwork()
	err := n.AddStation(&station{id: 1})
	if err != nil {
		t.Error("Failed to add first station to network")
	}
	err = n.AddStation(&station{id: 2})
	if err != nil {
		t.Error("Failed to add second station to network")
	}
	_, err = n.GetStation(1)
	if err != nil {
		t.Error("Failed to get first station")
	}
	_, err = n.GetStation(2)
	if err != nil {
		t.Error("Failed to get second station")
	}
}

func TestReachFromNonExistSourceStation(t *testing.T) {
	n := NewNetwork()
	src := &station{id: 1}
	dst := &station{id: 2}

	n.AddStation(dst)
	reach, err := n.CheckReachability(src, dst)
	if reach || err == nil {
		t.Errorf("Should not be able to reach a source station that was not added to network")
	}
}

func TestReachFromNonExistDestinationStation(t *testing.T) {
	n := NewNetwork()
	src := &station{id: 1}
	dst := &station{id: 2}

	n.AddStation(src)
	reach, err := n.CheckReachability(src, dst)
	if err != nil {
		t.Errorf("Expected to checkReachability without errors %s", err)
	}
	if reach {
		t.Errorf("Expected to not reach without errors")
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
		t.Errorf("Expected to checkReachability without errors %s", err)
	}
	if reach {
		t.Errorf("Expected to not reach")
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
		t.Errorf("Expected to checkReachability without errors %s", err)
	}
	if !reach {
		t.Errorf("Expected to reach")
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
		t.Errorf("Unexpected error %s", err)
	}
	if !reach {
		t.Errorf("Expected to reach")
	}
}

func TestConnectBetweenStationsNotExist(t *testing.T) {
	n := NewNetwork()
	src := &station{id: 1}
	dst := &station{id: 2}
	err := n.ConnectStations(src, dst)
	if err == nil {
		t.Error("Expected not to connect between station that don't exist")
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
		t.Error("Expect to connect stations")
	}
	reach, err := n.CheckReachability(src, dst)
	if !reach || err != nil {
		t.Error("Expected to reach after connected")
	}
}
