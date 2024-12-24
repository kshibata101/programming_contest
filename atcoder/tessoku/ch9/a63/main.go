package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc *bufio.Scanner

func Scan() int {
	sc.Scan()
	a, _ := strconv.Atoi(sc.Text())
	return a
}

type P struct {
	V int
	D int
}

func main() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	n := Scan()
	m := Scan()
	g := make([][]int, n+1)
	for i := 0; i < n; i++ {
		g[i+1] = []int{}
	}
	for i := 0; i < m; i++ {
		a := Scan()
		b := Scan()
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	dis := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dis[i] = -1
	}
	q := []P{}
	q = append(q, P{1, 0})
	for len(q) > 0 {
		p := q[0]
		if len(q) > 1 {
			q = q[1:]
		} else {
			q = []P{}
		}
		if dis[p.V] == -1 || p.D < dis[p.V] {
			dis[p.V] = p.D
			for i := 0; i < len(g[p.V]); i++ {
				nei := g[p.V][i]
				q = append(q, P{nei, p.D + 1})
			}
		}
	}

	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	for i := 1; i <= n; i++ {
		fmt.Fprintln(wr, dis[i])
	}
}
