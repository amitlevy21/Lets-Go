package main

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
	id             int64
	vCategory      category
	manufactor     string
	manufactorDate time.Time
	activeSince    time.Time
	capacity       int
	seats          int
}
