package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var s *bufio.Scanner

func Scan() int {
	s.Scan()
	a, _ := strconv.Atoi(s.Text())
	return a
}

func main() {
	s = bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	n := Scan()
	q := Scan()
	dp := make([][]int, 30)
	dp[0] = make([]int, n+1)
	for i := 0; i < n; i++ {
		dp[0][i+1] = Scan()
	}
	for i := 1; i < 30; i++ {
		dp[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			dp[i][j] = dp[i-1][dp[i-1][j]]
		}
	}

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	for i := 0; i < q; i++ {
		x := Scan()
		y := Scan()
		for d := 0; d < 30; d++ {
			if (1<<d)&y != 0 {
				x = dp[d][x]
			}
		}
		fmt.Fprintln(w, x)
	}
}
