package main

import (
	"fmt"
	"math"
)

type Triangle struct {
	base   float64
	height float64
}

func (t Triangle) area() float64 {
	return math.Round((t.height * t.base) / 2)
}

func (t Triangle) perimiter() float64 {
	return math.Round((math.Sqrt(math.Pow(t.height, 2) + math.Pow(t.base, 2))) + t.height + t.base)
}

func main() {
	triangle := Triangle {
		base: 25, height: 29,
	}
	fmt.Println("Triangle area is: ", triangle.area())
	fmt.Println("Triangle perimiter is: ", triangle.perimiter())
}
