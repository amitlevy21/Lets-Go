package ride

import "time"

type cycle int8

const (
	hourly cycle = iota
	weekly
	yearly
)

type schedule struct {
	id                  int64
	stationStartID      int64
	stationFinishID     int64
	duration            time.Time
	leaveFromSource     time.Time
	arriveAtDestination time.Time
	repeats             cycle
}
