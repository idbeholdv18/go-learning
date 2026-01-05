package geometry

import "math"

type Point struct {
	X float64
	Y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-q.Y)
}
