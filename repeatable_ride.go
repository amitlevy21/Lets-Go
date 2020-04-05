package main

import "fmt"

type repeatableRide struct {
	id          uint64
	stationsIds []uint64
}

// NewRepeatable creates a new ride that its data will not change between
// ride executions. stationIds can be nil, however equally consecutive ids
// will return an error.
func NewRepeatable(stationIds []uint64) (*repeatableRide, error) {
	if stationIds == nil {
		return &repeatableRide{stationsIds: make([]uint64, 0)}, nil
	}
	for i, sID := range stationIds[1:] {
		last := stationIds[i]
		if err := checkStationToItself(last, sID); err != nil {
			return nil, err
		}
	}
	return &repeatableRide{stationsIds: stationIds}, nil
}

func (rr *repeatableRide) AddStation(stationID uint64) error {
	last := rr.stationsIds[len(rr.stationsIds)-1]
	if err := checkStationToItself(last, stationID); err != nil {
		return err
	}
	rr.stationsIds = append(rr.stationsIds, stationID)
	return nil
}

func checkStationToItself(src uint64, dst uint64) error {
	if src == dst {
		return fmt.Errorf("cannot go from station %d to itself", dst)
	}
	return nil
}
