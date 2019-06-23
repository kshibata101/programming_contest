package main

import (
	"fmt"
)

var dim int = 4
var n int = dim * dim

func main() {
	n := 16
	puzzle := make([]int, n)
	var zero int
	for i := 0; i < n; i++ {
		fmt.Scan(&puzzle[i])
		if puzzle[i] == 0 {
			puzzle[i] = n
			zero = i
		}
	}

	res := search(&puzzle, 0, zero, -100)
	fmt.Println(res)
}

func search(p *[]int, turn int, zero int, prev_dir int) int {
	dis := dist(p)
	if dis == 0 {
		return turn
	}
	dir := []int{-dim,-1,1,dim}
	for i := 0; i < len(dir); i++ {
		if dir[i] + prev_dir == 0 { // do not back to previous state
			continue
		}
		next := zero + dir[i]
		if next < 0 || n <= next {
			continue
		}
		if (next%dim - zero%dim) * dir[i] < 0 {
			continue
		}
		
		(*p)[next], (*p)[zero] = (*p)[zero], (*p)[next]
		res := search(p, turn+1, next, dir[i])
		(*p)[next], (*p)[zero] = (*p)[zero], (*p)[next]
		if res >= 0 {
			return res
		}
	}
	return -1
}

func dist(p *[]int) int {
	sum := 0
	for i := 0; i < n; i++ {
		v := (*p)[i] - 1
		sum += abs(v/dim - i/dim)
		sum += abs(v%dim - i%dim)
	}
	return sum
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
