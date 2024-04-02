package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var s *bufio.Scanner

func ScanI() int {
	s.Scan()
	a, _ := strconv.Atoi(s.Text())
	return a
}

func ScanIArr(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = ScanI()
	}
	return a
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func main() {
	s = bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)

	n := ScanI()
	m := ScanI()
	a := make([]int, m)
	dp := make([][]int, m+1)
	inf := 1 << 31
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			a[i] += ScanI() * (1 << j)
		}
	}

	for i := 0; i <= m; i++ {
		dp[i] = make([]int, 1<<n)
		for j := 0; j < 1<<n; j++ {
			dp[i][j] = inf
		}
	}
	dp[0][0] = 0

	for i := 0; i < m; i++ {
		for j := 0; j < 1<<n; j++ {
			if dp[i][j] == inf {
				continue
			}

			dp[i+1][j] = Min(dp[i+1][j], dp[i][j])
			dp[i+1][j|a[i]] = Min(dp[i+1][j|a[i]], dp[i][j]+1)
		}
	}
	ans := dp[m][1<<n-1]
	if ans == inf {
		ans = -1
	}
	fmt.Println(ans)
}
