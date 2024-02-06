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
	a := make([]int, n)
	a[0] = 10000000
	for i := 1; i < n; i++ {
		a[i] = load(s)
	}
	b := make([]int, n)
	b[0] = 10000000
	b[1] = 10000000
	for i := 2; i < n; i++ {
		b[i] = load(s)
	}

	dp := make([]int, n)
	for i := 1; i < n; i++ {
		dp[i] = dp[i-1] + a[i]
		if i >= 2 && dp[i-2]+b[i] < dp[i] {
			dp[i] = dp[i-2] + b[i]
		}
	}
	fmt.Println(dp[n-1])
}

func load(s *bufio.Scanner) int {
	s.Scan()
	a, _ := strconv.Atoi(s.Text())
	return a
}
