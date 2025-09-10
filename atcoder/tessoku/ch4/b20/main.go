package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc *bufio.Scanner

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

func main() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	s := ScanS()
	t := ScanS()
	dp := make([][]int, len(s)+1)
	for i := 0; i <= len(s); i++ {
		dp[i] = make([]int, len(t)+1)
	}
	for i := 0; i <= len(t); i++ {
		dp[0][i] = i
	}
	for i := 0; i <= len(s); i++ {
		dp[i][0] = i
	}
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(t); j++ {
			dp[i][j] = Min(dp[i-1][j]+1, dp[i][j-1]+1)
			c := 1
			if s[i-1] == t[j-1] {
				c = 0
			}
			dp[i][j] = Min(dp[i][j], dp[i-1][j-1]+c)
		}
	}
	fmt.Println(dp[len(s)][len(t)])
}
