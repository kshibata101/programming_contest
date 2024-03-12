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
	a := loadArr(s, n-1)
	b := loadArr(s, n-1)

	dp := make([]int, n)
	for i := 1; i < n; i++ {
		dp[i] = -(1 << 31)
	}
	for i := 0; i < n-1; i++ {
		dp[a[i]-1] = max(dp[a[i]-1], dp[i]+100)
		dp[b[i]-1] = max(dp[b[i]-1], dp[i]+150)
	}
	fmt.Println(dp[n-1])
}

func load(s *bufio.Scanner) int {
	s.Scan()
	a, _ := strconv.Atoi(s.Text())
	return a
}

func loadArr(s *bufio.Scanner, n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = load(s)
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
