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

	n := Scan()
	W := Scan()
	w := make([]int, n+1)
	v := make([]int, n+1)
	for i := 0; i < n; i++ {
		w[i+1] = Scan()
		v[i+1] = Scan()
	}

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, 100001)
		for j := 0; j <= 100000; j++ {
			dp[i][j] = W + 1
		}
	}
	dp[0][0] = 0
	for i := 1; i <= n; i++ {
		for j := 0; j <= 100000; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= v[i] {
				if dp[i-1][j-v[i]]+w[i] <= W {
					dp[i][j] = Min(dp[i][j], dp[i-1][j-v[i]]+w[i])
				}
			}
		}
	}
	ans := 0
	for i := 100000; i >= 0; i-- {
		if dp[n][i] <= W {
			ans = i
			break
		}
	}
	fmt.Println(ans)
}
