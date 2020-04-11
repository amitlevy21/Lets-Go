package ride

import (
	"fmt"
	"strings"
)

type route struct {
	id          int64
	stationsIds []int64
}

type stationAdder interface {
	AddStation(id int64) error
}

func (r *route) AddStation(stationID int64, a stationAdder) error {
	last := r.stationsIds[len(r.stationsIds)-1]
	if err := checkStationToItself(last, stationID); err != nil {
		return err
	}
	// TODO: create error handling types
	// BODY: use types like stationExist for more readable and relyable checks
	if err := a.AddStation(stationID); err != nil && !strings.Contains(err.Error(), "exist") {
		return err
	}
	r.stationsIds = append(r.stationsIds, stationID)
	return nil
}

func checkStationToItself(src, dst int64) error {
	if src == dst {
		return fmt.Errorf("cannot go from station %d to itself", dst)
	}
	return nil
}
