package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

type Circle struct {
	Point
	radius float64
}

func main() {
	var circle1 Circle = Circle{
		Point: Point{
			x: 10,
			y: 7,
		},
		radius: 4,
	}
	var circle2 Circle = Circle{
		Point: Point{
			x: 1,
			y: 2,
		},
		radius: 2,
	}

	fmt.Println(isOverflow(&circle1, &circle2))
}

func distance(a *Point, b *Point) float64 {
	x := math.Pow(float64(a.x)-float64(b.x), 2)
	y := math.Pow(float64(a.y)-float64(b.y), 2)
	return math.Sqrt(x + y)
}

func isOverflow(a *Circle, b *Circle) bool {
	if distance(&a.Point, &b.Point) < a.radius+b.radius {
		return true
	}
	return false
}
