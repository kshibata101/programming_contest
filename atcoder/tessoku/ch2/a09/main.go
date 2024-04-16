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
	h := load(s)
	w := load(s)
	n := load(s)
	m := make([][]int, h+1)
	ans := make([][]int, h+1)
	for i := 0; i <= h; i++ {
		m[i] = make([]int, w+1)
		ans[i] = make([]int, w+1)
	}

	for i := 0; i < n; i++ {
		a := load(s)
		b := load(s)
		c := load(s)
		d := load(s)
		m[a][b]++
		if c < h {
			m[c+1][b]--
		}
		if d < w {
			m[a][d+1]--
		}
		if c < h && d < w {
			m[c+1][d+1]++
		}
	}

	for i := 1; i <= h; i++ {
		for j := 1; j <= w; j++ {
			ans[i][j] = ans[i-1][j] + ans[i][j-1] - ans[i-1][j-1] + m[i][j]
		}
	}
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()
	for i := 1; i <= h; i++ {
		fmt.Fprint(wr, ans[i][1])
		for j := 2; j <= w; j++ {
			fmt.Fprint(wr, " ")
			fmt.Fprint(wr, ans[i][j])
		}
		fmt.Fprintln(wr)
	}
}

func load(s *bufio.Scanner) int {
	s.Scan()
	a, _ := strconv.Atoi(s.Text())
	return a
}
