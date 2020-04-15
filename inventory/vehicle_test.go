// Copyright (c) 2020 Amit Levy
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package inventory

import (
	"testing"
	"time"
)

func TestCreateVehicle(t *testing.T) {
	v := vehicle{
		vCategory:      bus,
		manufactor:     "Toyota",
		manufactorDate: time.Date(1999, time.April, 1, 19, 30, 50, 10, time.UTC),
		activeSince:    time.Date(2000, time.April, 1, 19, 30, 50, 10, time.UTC),
		capacity:       30,
		seats:          10,
	}
	t.Logf("Vehicle created! %v", v)
}
