package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

func (p Point) Distance(other Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	a := NewPoint(3, 4)
	b := NewPoint(0, 0)
	distance := a.Distance(b)

	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}
