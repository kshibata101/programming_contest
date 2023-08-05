package main

import "fmt"

type pond struct {
	x int
	a int
}

func main() {
	var s string
	fmt.Scan(&s)
	var x, area int

	var ss []int = make([]int, 0)
	var ps []*pond = make([]*pond, 0)

	for i := 0; i < len(s); i++ {
		x++
		switch string(s[i]) {
		case "/":
			y := ss[len(ss)-1]
			ss = ss[:len(ss)-1]

			a := x - y - 1
			for len(ps) > 0 {
				p := ps[len(ps)-1]
				if p.x <= y {
					break
				}
				a += p.a
				ps = ps[:len(ps)-1]
			}
			ps = append(ps, &pond{a: a, x: y})
		case "\\":
			ss = append(ss, x-1)
		}
	}

	fmt.Println(area)
	fmt.Print(len())
	for i := 0; i < len(ps); i++ {
		if ps[i].a > 0 {
			fmt.Printf(" %d", ps[i].a)
		}
	}
	fmt.Println()
}
