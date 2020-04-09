package ride

import "fmt"

type repeatableRide struct {
	id          int64
	stationsIds []int64
}

func (rr *repeatableRide) AddStation(stationID int64) error {
	last := rr.stationsIds[len(rr.stationsIds)-1]
	if err := checkStationToItself(last, stationID); err != nil {
		return err
	}
	rr.stationsIds = append(rr.stationsIds, stationID)
	return nil
}

func checkStationToItself(src int64, dst int64) error {
	if src == dst {
		return fmt.Errorf("cannot go from station %d to itself", dst)
	}
	return nil
}
