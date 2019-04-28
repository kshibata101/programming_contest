package main

import (
	"fmt"
)

type Pod struct {
	Pos int
	Area int
}

func main() {
	var str string
	fmt.Scan(&str)

	downs := []int{}
	pods := []Pod{}
	sum := 0
	
	rs := []rune(str)
	for i := 0; i < len(rs); i++ {
		s := string(rs[i])
		switch s {
		case "\\":
			downs = append(downs, i)
		case "/":
			if len(downs) <= 0 {
				continue
			}
			down := downs[len(downs)-1]
			downs = downs[0:len(downs)-1]

			area := i - down
			sum += area

			for j := len(pods) - 1; j >= 0; j-- {
				pod := pods[j]
				if down < pod.Pos {
					area += pod.Area
					pods = pods[0:j]
				} else {
					break
				}
			}
			pods = append(pods, Pod{down, area})
		}
	}

	fmt.Println(sum)
	fmt.Print(len(pods))
	if len(pods) > 0 {
		for i := 0; i < len(pods); i++ {
			fmt.Print(" ")
			fmt.Print(pods[i].Area)
		}
	}
	fmt.Println()
}

