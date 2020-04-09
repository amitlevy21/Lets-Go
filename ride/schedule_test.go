package ride

import "testing"

func TestCreateSchedule(t *testing.T) {
	s := schedule{
		stationStartID:  int64(0),
		stationFinishID: int64(1),
		repeats:         hourly,
	}
	t.Logf("schedule created! %v", s)
}
