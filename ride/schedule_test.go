// Copyright (c) 2020 Amit Levy
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package ride

import (
	"testing"
	"time"
)

func TestCreateSchedule(t *testing.T) {
	s := schedule{
		leaveSrc:  time.Date(1991, time.April, 3, 10, 30, 0, 0, time.UTC),
		arriveDst: time.Date(1991, time.April, 3, 11, 20, 0, 0, time.UTC),
		duration:  50 * time.Minute,
	}
	t.Logf("schedule created! %v", s)
}
