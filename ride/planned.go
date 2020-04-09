package ride

import "fmt"

type status int8

const (
	pending status = iota
	ongoing
	finished
	cancelled
)

var statusStr = []string{"pending", "ongoing", "finished", "cancelled"}

type planned struct {
	id             uint64
	rStatus        status
	numPassengers  int32
	availableSeats int32
	crewMembersIds []uint64
	vehicle        int64
	latestStation  int64
}

func (or *planned) start() error {
	if or.rStatus != pending {
		return fmt.Errorf("ride already started")
	}
	or.rStatus = ongoing
	return nil
}

func (or *planned) finish() error {
	if or.rStatus == finished {
		return fmt.Errorf("ride already finished")
	}
	if or.rStatus == cancelled {
		return fmt.Errorf("cannot finish cancelled ride")
	}
	or.rStatus = finished
	return nil
}

func (or *planned) cancel() error {
	or.rStatus = cancelled
	return nil
}

func (s status) String() string {
	return statusStr[s]
}
