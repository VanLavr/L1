package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func New(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func (p Point) FindDistance(anotherPoint Point) float64 {
	return math.Sqrt(math.Pow(p.x-anotherPoint.x, 2) + math.Pow(p.y-anotherPoint.y, 2))
}

func main() {
	p1 := New(-2, 1)
	p2 := New(2, -2)

	fmt.Println(p1.FindDistance(*p2))
}
