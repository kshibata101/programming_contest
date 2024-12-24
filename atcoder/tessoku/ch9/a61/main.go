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

	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	for i := 1; i <= n; i++ {
		fmt.Fprint(wr, i, ": {")
		for j := 0; j < len(g[i]); j++ {
			fmt.Fprint(wr, g[i][j])
			if j < len(g[i])-1 {
				fmt.Fprint(wr, ", ")
			}
		}
		fmt.Fprintln(wr, "}")
	}
}
