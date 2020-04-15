// Copyright (c) 2020 Amit Levy
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package ride

import "time"

type schedule struct {
	id        int64
	duration  time.Duration
	leaveSrc  time.Time
	arriveDst time.Time
}
