package main

import "time"

type person struct {
	firstName string
	lastName string
	age int16
}

type role int

const (
	driver role = iota
	engineer
	captain
)

type crewMember struct {
	person
	cRole role
	firstWorkDay time.Time
}