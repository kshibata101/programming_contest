package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)

	n := load(s)
	g := make([][]int, n+1)
	for i := 1; i < n; i++ {
		u := load(s)
		v := load(s)
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}
	sum := 1
	max := 0
	for i := 0; i < len(g[1]); i++ {
		cost := search(g, g[1][i], 1)
		if cost > max {
			max = cost
		}
		sum += cost
	}
	fmt.Println(sum - max)
}

func search(g [][]int, v, from int) int {
	sum := 1
	for i := 0; i < len(g[v]); i++ {
		if g[v][i] == from {
			continue
		}
		sum += search(g, g[v][i], v)
	}
	return sum
}

func load(s *bufio.Scanner) int {
	s.Scan()
	a, _ := strconv.Atoi(s.Text())
	return a
}