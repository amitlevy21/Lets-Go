package main

import "testing"

func TestCreateSchedule(t * testing.T) {
	s := schedule{
		source: station{},
		destination: station{},
		repeats: hourly,
	}

	t.Logf("schedule created! %v", s)
}