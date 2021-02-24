package main

import "fmt"

type pond struct {
	start int
	area  int
}

func main() {
	var s string
	fmt.Scan(&s)
	ls := []rune(s)

	ss := make([]int, 0)
	ps := make([]*pond, 0)

	sum := 0
	for i := 0; i < len(ls); i++ {
		str := (string)(ls[i])
		if str == "\\" {
			ss = append(ss, i)
		} else if str == "/" {
			if len(ss) == 0 {
				continue
			}
			start := ss[len(ss)-1]
			ss = ss[:len(ss)-1]

			area := i - start
			sum += area

			for j := len(ps) - 1; j >= 0; j-- {
				if ps[j].start <= start {
					break
				}
				area += ps[j].area
				ps = ps[:j]
			}
			ps = append(ps, &pond{start: start, area: area})
		}
	}

	fmt.Println(sum)
	fmt.Print(len(ps))
	for i := 0; i < len(ps); i++ {
		fmt.Printf(" %d", ps[i].area)
	}
	fmt.Println()
}
