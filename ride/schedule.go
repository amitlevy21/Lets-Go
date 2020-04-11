package ride

import "time"

type schedule struct {
	id        int64
	duration  time.Duration
	leaveSrc  time.Time
	arriveDst time.Time
}
