package main

import (
	"fmt"
)

type Pos struct {
	X int
	Y int
}

type Area struct {
	X1 int
	X2 int
	Y int
	Area int
}

func main() {
	var str string
	fmt.Scan(&str)

	poses := []Pos{Pos{0, 0}}
	areas := []Area{}
	
	rs := []rune(str)
	y := 0
	for x := 1; x <= len(rs); x++ {
		s := string(rs[x-1])
		switch s {
		case "\\":
			y--
		case "/":
			y++
			under_area := 0
			delete_area_index := -1
			for i := len(poses) - 1; i >= 0; i-- {
				pos := poses[i]
				if pos.Y < y {
					// areaをチェックして、いたら消す
					for j := len(areas) - 1; j >= 0; j-- {
						if pos.X == areas[j].X2 {
							under_area += areas[j].Area
							delete_area_index = j
							break
						}
					}
				} else if pos.Y == y {
					if delete_area_index >= 0 {
						areas = areas[0:delete_area_index]
					}
					poses = poses[0:i+1]
					areas = append(areas, Area{pos.X, x, y, under_area + x - pos.X - 1})
					break
				} else {
					// ないはず
					break
				}
			}
		}
		poses = append(poses, Pos{x, y})
	}

	a := 0
	k := len(areas)
	for i := 0; i < k; i++ {
		a += areas[i].Area
	}
	fmt.Println(a)
	fmt.Print(k)
	if len(areas) > 0 {
		fmt.Print(" ")
		fmt.Print(areas[0].Area)
		for i := 1; i < k; i++ {
			fmt.Print(" ")
			fmt.Print(areas[i].Area)
		}
	}
	fmt.Println()
}
