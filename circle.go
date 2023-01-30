package main

import (
	"math"
)

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Round(math.Pi * math.Pow(c.radius, 2))
}

func (c Circle) perimiter() float64 {
	return math.Round(2 * math.Pi * c.radius)
}
