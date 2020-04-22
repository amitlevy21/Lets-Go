// Copyright (c) 2020 Amit Levy
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package ride

import (
	"fmt"
	"strings"
)

// Route is the order a ride visits stations.
type Route struct {
	id          int64
	StationsIds []int64
	Schedules   []Schedule
}

type stationAdder interface {
	Add(id int64) error
}

// AddStation adds a station to the route.
func (r *Route) AddStation(stationID int64, a stationAdder) error {
	last := r.StationsIds[len(r.StationsIds)-1]
	if err := checkStationToItself(last, stationID); err != nil {
		return err
	}
	// TODO create error handling types
	// BODY use types like stationExist for more readable and relyable checks
	if err := a.Add(stationID); err != nil && !strings.Contains(err.Error(), "exist") {
		return err
	}
	r.StationsIds = append(r.StationsIds, stationID)
	return nil
}

func checkStationToItself(src, dst int64) error {
	if src == dst {
		return fmt.Errorf("cannot go from station %d to itself", dst)
	}
	return nil
}
