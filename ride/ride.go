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

type ride struct {
	id             int64
	rStatus        status
	numPassengers  int32
	availableSeats int32
	crewMembersIds []int64
	vehicle        int64
	latestStation  int64
}

func (r *ride) start() error {
	if r.rStatus != pending {
		return fmt.Errorf("ride already started")
	}
	r.rStatus = ongoing
	return nil
}

func (r *ride) finish() error {
	if r.rStatus == finished {
		return fmt.Errorf("ride already finished")
	}
	if r.rStatus == cancelled {
		return fmt.Errorf("cannot finish cancelled ride")
	}
	r.rStatus = finished
	return nil
}

func (r *ride) cancel() error {
	r.rStatus = cancelled
	return nil
}

func (s status) String() string {
	return statusStr[s]
}
