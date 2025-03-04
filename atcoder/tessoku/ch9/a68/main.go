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

func Min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

type Edge struct {
	to  int
	cap int
	rev int
}

func Dfs(pos, goal, F int, used *[]bool, G [][]Edge) int {
	if pos == goal {
		return F
	}
	(*used)[pos] = true
	for j := 0; j < len(G[pos]); j++ {
		if G[pos][j].cap == 0 || (*used)[G[pos][j].to] {
			continue
		}
		f := Dfs(G[pos][j].to, goal, Min(F, G[pos][j].cap), used, G)
		if f > 0 {
			G[pos][j].cap -= f
			G[G[pos][j].to][G[pos][j].rev].cap += f
			return f
		}
	}
	return 0
}

func main() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	n := Scan()
	G := make([][]Edge, n+1)
	for j := 1; j <= n; j++ {
		G[j] = []Edge{}
	}

	m := Scan()
	for j := 0; j < m; j++ {
		a := Scan()
		b := Scan()
		c := Scan()
		as := len(G[a])
		bs := len(G[b])
		G[a] = append(G[a], Edge{b, c, bs})
		G[b] = append(G[b], Edge{a, 0, as})
	}
	ans := 0
	for {
		used := make([]bool, n+1)
		f := Dfs(1, n, 5001, &used, G)
		if f > 0 {
			ans += f
		} else {
			break
		}
	}
	fmt.Println(ans)
}
