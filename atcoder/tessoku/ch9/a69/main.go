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

func ScanS() string {
	sc.Scan()
	return sc.Text()
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
	G := make([][]Edge, 2*n+2)
	for j := 0; j < 2*n+2; j++ {
		G[j] = []Edge{}
	}

	for j := 1; j <= n; j++ {
		jl := len(G[j])
		njl := len(G[n+j])
		G[0] = append(G[0], Edge{j, 1, jl})
		G[j] = append(G[j], Edge{0, 0, j - 1})
		G[n+j] = append(G[n+j], Edge{2*n + 1, 1, j - 1})
		G[2*n+1] = append(G[2*n+1], Edge{n + j, 0, njl})

		s := ScanS()
		for k := 1; k <= len(s); k++ {
			if s[k-1] != '#' {
				continue
			}
			js := len(G[j])
			ks := len(G[n+k])
			G[j] = append(G[j], Edge{n + k, 1, ks})
			G[n+k] = append(G[n+k], Edge{j, 0, js})
		}
	}

	ans := 0
	for {
		used := make([]bool, 2*n+2)
		f := Dfs(0, 2*n+1, n, &used, G)
		if f > 0 {
			ans += f
		} else {
			break
		}
	}
	fmt.Println(ans)
}
