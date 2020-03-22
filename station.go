package main

type coordinates struct {
	x float64
	y float64
}

type station struct {
	location coordinates
	rideIds  []int64
}
