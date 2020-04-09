package ride

import "fmt"

type route struct {
	id          int64
	stationsIds []int64
}

func (r *route) AddStation(stationID int64) error {
	last := r.stationsIds[len(r.stationsIds)-1]
	if err := checkStationToItself(last, stationID); err != nil {
		return err
	}
	r.stationsIds = append(r.stationsIds, stationID)
	return nil
}

func checkStationToItself(src int64, dst int64) error {
	if src == dst {
		return fmt.Errorf("cannot go from station %d to itself", dst)
	}
	return nil
}
