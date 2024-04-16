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
	W := load(s)
	w := make([]int, n)
	v := make([]int, n)
	for i := 0; i < n; i++ {
		w[i] = load(s)
		v[i] = load(s)
	}

	dp := make([][]int, n+1)
	dp[0] = make([]int, W+1)
	for i := 0; i < n; i++ {
		dp[i+1] = make([]int, W+1)
		for j := 0; j <= W; j++ {
			dp[i+1][j] = max(dp[i+1][j], dp[i][j])
			if j+w[i] <= W {
				dp[i+1][j+w[i]] = dp[i][j] + v[i]
			}
		}
	}
	m := 0
	for i := 0; i <= W; i++ {
		m = max(m, dp[n][i])
	}
	fmt.Println(m)
}

func load(s *bufio.Scanner) int {
	s.Scan()
	a, _ := strconv.Atoi(s.Text())
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
