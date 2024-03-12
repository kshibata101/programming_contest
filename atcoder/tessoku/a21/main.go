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
	p := make([]int, n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = load(s) - 1
		a[i] = load(s)
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < n-1; i++ {
		for j := n - 1; j > i; j-- {
			// 左から消す場合
			dp[i+1][j] = dp[i][j]
			if i <= p[i] && p[i] <= j {
				dp[i+1][j] = dp[i][j] + a[i]
			}
			// 右から消す場合
			dp[i][j-1] = max(dp[i][j-1], dp[i][j])
			if i <= p[j] && p[j] <= j {
				dp[i][j-1] = dp[i][j] + a[j]
			}
		}
	}
	mx := 0
	for i := 0; i < n; i++ {
		mx = max(mx, dp[i][i])
	}
	fmt.Println(mx)
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
