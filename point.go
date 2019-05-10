package main

import "math"

type Point struct {
	X    int
	Y    int
	Cost int
	Parent *Point
}

func NewPoint(x, y int) *Point {
	return &Point{X: x, Y: y, Cost: math.MaxInt64}
}
