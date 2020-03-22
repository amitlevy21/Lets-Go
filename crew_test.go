package main

import (
	"testing"
	"time"
)

func TestCreateCrew(t *testing.T) {
	c := crewMember{
		person: person{
			firstName: "Dan",
			lastName:  "Pill",
			age:       28,
		},
		cRole: driver,
		firstWorkDay: time.Date(2009, time.April, 1, 8, 0, 0, 0, time.UTC),
	}

	t.Logf("Created crew! %v", c)
}
