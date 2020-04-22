// Copyright (c) 2020 Amit Levy
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package inventory

import "time"

type category int

const (
	bus category = iota
	ship
	taxi
	train
	metro
)

type vehicle struct {
	id             uint64
	vCategory      category
	manufactor     string
	manufactorDate time.Time
	activeSince    time.Time
	capacity       int
	seats          int
}
