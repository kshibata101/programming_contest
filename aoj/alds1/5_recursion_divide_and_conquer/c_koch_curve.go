package main

import (
	"fmt"
	"math"
)

type Point struct {
	X float64
	Y float64
}

func split(s, e Point) (Point, Point, Point) {
	rad60 := math.Pi / 3.0

	p1 := Point{(s.X * 2.0 + e.X) / 3.0, (s.Y * 2.0 + e.Y) / 3.0}
	p3 := Point{(s.X + e.X * 2.0) / 3.0, (s.Y + e.Y * 2.0) / 3.0}
	
	x, y := p3.X - p1.X, p3.Y - p1.Y
	rx := math.Cos(rad60) * x - math.Sin(rad60) * y
	ry := math.Sin(rad60) * x + math.Cos(rad60) * y

	p2 := Point{p1.X + rx, p1.Y + ry}

	return p1, p2, p3
}

func main() {
	points := []Point{Point{0.0, 0.0}, Point{100.0, 0.0}}

	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		lineCount := len(points) - 1
		newPoints := make([]Point, len(points) + lineCount * 3)
		p := 0
		for j := 0; j < lineCount; j++ {
			newPoints[p] = points[j]
			newPoints[p+1], newPoints[p+2], newPoints[p+3] = split(points[j], points[j+1])
			p += 4
		}
		newPoints[p] = points[len(points) - 1]
		points = newPoints
	}

	for i := 0; i < len(points); i++ {
		fmt.Printf("%.8f %.8f\n", points[i].X, points[i].Y)
	}
}
