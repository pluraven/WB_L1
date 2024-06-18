package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (p *Point) New() *Point {
	return p
}

func (p *Point) GetCoords() (float64, float64) {
	return p.x, p.y
}

func (p1 *Point) GetDistance(p2 *Point) float64 {
	x1, y1 := p1.GetCoords()
	x2, y2 := p2.GetCoords()
	return math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2))
}

func main() {
	a := Point{x: 1, y: 2}
	b := Point{x: 1, y: 1}
	fmt.Println("Расстояние между точками a и b", a.GetDistance(&b))
}
