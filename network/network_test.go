// Copyright (c) 2020 Amit Levy
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package network

import "testing"

func TestCreateNetwork(t *testing.T) {
	n := NewNetwork()
	t.Logf("network created! %v", n)
}

func TestGetNonExist(t *testing.T) {
	n := NewNetwork()
	if _, ok := n.Get(1); ok {
		t.Error("cannot get an element that doesn't exist in the network")
	}
}

func TestAddDuplicate(t *testing.T) {
	n := NewNetwork()
	e := int64(1)
	failIfError(t, n.Add(e))
	if err := n.Add(e); err == nil {
		t.Errorf("cannot add duplicate element to same network")
	}
}

func TestAdd(t *testing.T) {
	n := NewNetwork()
	s := int64(0)
	if err := n.Add(s); err != nil {
		t.Errorf("error while adding to network: %s", err)
	}
	if _, ok := n.Get(s); !ok {
		t.Errorf("failed to get added %d", s)
	}
}

func TestAddMultiple(t *testing.T) {
	n := NewNetwork()
	e1 := int64(0)
	e2 := int64(1)
	if err := n.Add(e1); err != nil {
		t.Error("failed to add first to network")
	}
	if err := n.Add(e2); err != nil {
		t.Error("failed to add second to network")
	}
	if _, ok := n.Get(e1); !ok {
		t.Error("failed to get first")
	}
	if _, ok := n.Get(e2); !ok {
		t.Error("failed to get second")
	}
}

func TestReachNonExistSrc(t *testing.T) {
	n := NewNetwork()
	failIfError(t, n.Add(2))
	if reach, err := n.CheckReachability(1, 2); reach || err == nil {
		t.Errorf("should not reach when source not in network")
	}
}

func TestReachNonExistDst(t *testing.T) {
	n := NewNetwork()
	failIfError(t, n.Add(1))
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
	failIfError(t, n.Add(src))
	failIfError(t, n.Add(dst))
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
	failIfError(t, n.Add(src))
	failIfError(t, n.Add(dst))
	failIfError(t, n.Connect(src, dst, 0))
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

	failIfError(t, n.Add(src))
	failIfError(t, n.Add(mid))
	failIfError(t, n.Add(dst))
	failIfError(t, n.Connect(src, mid, 0))
	failIfError(t, n.Connect(mid, dst, 0))

	reach, err := n.CheckReachability(src, dst)
	failIfError(t, err)
	if !reach {
		t.Errorf("expected to reach")
	}
}

func TestConnectBothNotExist(t *testing.T) {
	n := NewNetwork()
	src := int64(1)
	dst := int64(2)
	if err := n.Connect(src, dst, 0); err == nil {
		t.Error("expected not to connect when both not exist")
	}
}

func TestConnectToItself(t *testing.T) {
	n := NewNetwork()
	src := int64(1)
	failIfError(t, n.Add(src))
	if err := n.Connect(src, src, 0); err == nil {
		t.Error("expected not to connect to itself")
	}
}

func TestConnectNegDuration(t *testing.T) {
	n := NewNetwork()
	s1 := int64(1)
	s2 := int64(2)
	failIfError(t, n.Add(s1))
	failIfError(t, n.Add(s2))
	if err := n.Connect(s1, s2, -1); err == nil {
		t.Error("should not accept negative duration")
	}
}

func TestConnect(t *testing.T) {
	n := NewNetwork()
	src := int64(1)
	dst := int64(2)
	failIfError(t, n.Add(src))
	failIfError(t, n.Add(dst))
	failIfError(t, n.Connect(src, dst, 1))
	reach, err := n.CheckReachability(src, dst)
	failIfError(t, err)
	if !reach {
		t.Error("expected to reach after connected")
	}
}

func TestValidateEmptyPath(t *testing.T) {
	n := NewNetwork()
	route := []int64{}
	if valid, _ := n.ValidatePath(route); valid {
		t.Errorf("should not be valid")
	}
}

func TestValidatePathStartNotExist(t *testing.T) {
	n := NewNetwork()
	route := []int64{1}
	if valid, err := n.ValidatePath(route); valid || err == nil {
		t.Errorf("should return error stating the route is invalid")
	}
}

func TestValidateUnreachablePath(t *testing.T) {
	n := NewNetwork()
	s1 := int64(1)
	s2 := int64(2)
	failIfError(t, n.Add(s1))
	failIfError(t, n.Add(s2))
	route := []int64{s1, s2}
	if valid, _ := n.ValidatePath(route); valid {
		t.Errorf("expected route not to be valid")
	}
}

func TestValidatePath(t *testing.T) {
	n := NewNetwork()
	s1 := int64(1)
	s2 := int64(2)
	s3 := int64(3)
	s4 := int64(4)
	failIfError(t, n.Add(s1))
	failIfError(t, n.Add(s2))
	failIfError(t, n.Add(s3))
	failIfError(t, n.Add(s4))
	failIfError(t, n.Connect(s1, s2, 0))
	failIfError(t, n.Connect(s2, s3, 0))
	failIfError(t, n.Connect(s2, s4, 0))
	route := []int64{s1, s2, s3}
	if valid, _ := n.ValidatePath(route); !valid {
		t.Errorf("expected route to be valid")
	}
}

func TestValidateDirectPath(t *testing.T) {
	n := NewNetwork()
	s1 := int64(1)
	s2 := int64(2)
	s3 := int64(3)
	s4 := int64(4)
	failIfError(t, n.Add(s1))
	failIfError(t, n.Add(s2))
	failIfError(t, n.Add(s3))
	failIfError(t, n.Add(s4))
	failIfError(t, n.Connect(s1, s2, 0))
	failIfError(t, n.Connect(s2, s3, 0))
	failIfError(t, n.Connect(s2, s4, 0))
	route := []int64{s1, s2, s4}
	if valid, _ := n.ValidatePath(route); !valid {
		t.Errorf("expected route to be valid")
	}
}

func failIfError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("unexpected error during test: %s", err)
	}
}
