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

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	n := Scan()
	s := ScanS()
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
		if i != n {
			dp[i][i+1] = 1
		}
	}
	for d := 2; d <= n; d++ {
		for l := 0; l+d <= n; l++ {
			r := l + d
			dp[l][r] = Max(dp[l+1][r], dp[l][r-1])
			if s[l] == s[r-1] {
				dp[l][r] = Max(dp[l][r], dp[l+1][r-1]+2)
			}
		}
	}
	fmt.Println(dp[0][n])
}
