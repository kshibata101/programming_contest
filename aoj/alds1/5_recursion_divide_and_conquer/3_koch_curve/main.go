package main

import (
	"fmt"
	"math"
)

type point [2]float64

var r3 = math.Sqrt(3)

func main() {
	var n int
	fmt.Scan(&n)

	var p1 point = [2]float64{0, 0}
	var p2 point = [2]float64{100, 0}
	points := []*point{&p1, &p2}

	for i := 0; i < n; i++ {
		news := make([]*point, len(points)*4-3)
		for j := 0; j < len(points)-1; j++ {
			po1, po2 := points[j], points[j+1]
			po3, po4, po5 := calcMidPoints(po1, po2)
			news[j*4] = po1
			news[j*4+1] = po3
			news[j*4+2] = po4
			news[j*4+3] = po5
			news[j*4+4] = po2
		}
		points = news
	}

	for i := 0; i < len(points); i++ {
		p := points[i]
		fmt.Printf("%.8f %.8f\n", p[0], p[1])
	}
}

func calcMidPoints(p1, p2 *point) (*point, *point, *point) {
	var p3 point = [2]float64{
		(2*p1[0] + p2[0]) / 3,
		(2*p1[1] + p2[1]) / 3,
	}
	var p5 point = [2]float64{
		(p1[0] + 2*p2[0]) / 3,
		(p1[1] + 2*p2[1]) / 3,
	}
	var p4 point = [2]float64{
		(3*(p1[0]+p2[0]) - r3*(p2[1]-p1[1])) / 6,
		(r3*(p2[0]-p1[0]) + 3*(p1[1]+p2[1])) / 6,
	}
	return &p3, &p4, &p5
}
