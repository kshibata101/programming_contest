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

func Dfs(g [][]int, visited []bool, i int) {
	if visited[i] {
		return
	}
	visited[i] = true
	for j := 0; j < len(g[i]); j++ {
		Dfs(g, visited, g[i][j])
	}
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
	visited := make([]bool, n+1)
	Dfs(g, visited, 1)

	ans := "The graph is connected."
	for i := 1; i <= n; i++ {
		if !visited[i] {
			ans = "The graph is not connected. "
			break
		}
	}
	fmt.Println(ans)
}
