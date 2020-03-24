package main

type coordinates struct {
	x float64
	y float64
}

type station struct {
	id       int64
	location coordinates
	rideIds  []int64
}
