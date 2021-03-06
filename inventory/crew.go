// Copyright (c) 2020 Amit Levy
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package inventory

import "time"

type person struct {
	firstName string
	lastName  string
	age       int16
}

type role int

const (
	driver role = iota
	engineer
	captain
)

type crewMember struct {
	person
	id           uint64
	cRole        role
	firstWorkDay time.Time
}
